package main

import (
	"context"
	"io"

	"github.com/Cacsjep/acapgrpcclient/cctv"
	"github.com/Cacsjep/goxis/pkg/acapapp"
	"github.com/Cacsjep/goxis/pkg/axvdo"
	grpc "google.golang.org/grpc"
)

var (
	// SERVER_ADDRESS is the address of the gRPC server.
	SERVER_ADDRESS = "10.0.0.54:50051"
	// vdo_format defines the video format as H264.
	vdo_format = axvdo.VdoFormatH264
	// stream_cfg sets up the video stream configuration with the specified format.
	stream_cfg = axvdo.VideoSteamConfiguration{Format: &vdo_format}
)

func main() {
	// Create a new ACAP application instance.
	app := acapapp.NewAcapApplication()
	app.Syslog.Info("Hello from " + app.Manifest.ACAPPackageConf.Setup.AppName + "!")

	// Create a new FrameProvider with the specified stream configuration.
	if err := app.NewFrameProvider(stream_cfg); err != nil {
		app.Syslog.Crit(err.Error())
	}

	// Start the FrameProvider to begin receiving frames.
	if err := app.FrameProvider.Start(); err != nil {
		app.Syslog.Crit(err.Error())
	}

	// Establish a connection to the gRPC server.
	conn, err := grpc.NewClient(SERVER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		app.Syslog.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new CCTV service client.
	client := cctv.NewCCTVServiceClient(conn)
	// Create a stream to send video frames.
	stream, err := client.StreamVideo(context.Background())
	if err != nil {
		app.Syslog.Errorf("error creating grpc stream: %v", err)
	}

	// Get the statistics of the FrameProvider.
	stats, err := app.FrameProvider.Stats()
	if err != nil {
		app.Syslog.Errorf("error getting stats: %v", err)
	}

loop:
	for {
		select {
		// Wait for a frame from the FrameProvider.
		case frame := <-app.FrameProvider.FrameStreamChannel:

			// Check if the frame has an error.
			if frame.Error != nil {
				app.Syslog.Errorf("Unexpected Vdo Error: %s", frame.Error.Error())
				continue
			}
			app.Syslog.Info(frame.String())

			// Create a frame message to send to the server.
			frame_message := &cctv.FrameMessage{
				Frame:       frame.Data,
				Height:      int32(stats.StreamStats.Height),
				Width:       int32(stats.StreamStats.Width),
				SequenceNbr: int32(frame.SequenceNbr),
				Framerate:   int32(stats.StreamStats.Framerate),
			}
			// Send the frame message to the server.
			if err := stream.Send(frame_message); err != nil {
				app.Syslog.Errorf("error sending frame: %v", err)
				if err == io.EOF {
					break loop
				}
				continue
			}

			// Receive acknowledgment from the server.
			_, err := stream.Recv()
			if err == io.EOF {
				app.Syslog.Info("Stream closed by server.")
				break loop
			}
			if err != nil {
				app.Syslog.Errorf("error receiving acknowledgment: %v", err)
				continue
			}
		}
	}

	// Close the stream when done.
	if err := stream.CloseSend(); err != nil {
		app.Syslog.Errorf("error closing stream: %v", err)
	}
}
