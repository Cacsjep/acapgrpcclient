package main

import (
	context "context"
	"io"

	"github.com/Cacsjep/acapgrpcclient/cctv"
	"github.com/Cacsjep/goxis/pkg/acapapp"
	"github.com/Cacsjep/goxis/pkg/axvdo"
	grpc "google.golang.org/grpc"
)

var (
	SERVER_ADDRESS = "10.0.0.62:50051"
	vdo_format     = axvdo.VdoFormatH264
	stream_cfg     = axvdo.VideoSteamConfiguration{Format: &vdo_format}
)

func main() {

	app := acapapp.NewAcapApplication()
	app.Syslog.Info("Hello from " + app.Manifest.ACAPPackageConf.Setup.AppName + "!")

	// FrameProvider for easy go channeld based frame receiving
	if err := app.NewFrameProvider(stream_cfg); err != nil {
		app.Syslog.Crit(err.Error())
	}

	// Start the frameprovider
	if err := app.FrameProvider.Start(); err != nil {
		app.Syslog.Crit(err.Error())
	}

	conn, err := grpc.NewClient(SERVER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		app.Syslog.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := cctv.NewCCTVServiceClient(conn)
	stream, err := client.StreamVideo(context.Background())
	if err != nil {
		app.Syslog.Errorf("error creating grpc stream: %v", err)
	}

loop:
	for {
		select {
		// Wait for frame from FrameProvider
		case frame := <-app.FrameProvider.FrameStreamChannel:

			// Check if frame has error
			if frame.Error != nil {
				app.Syslog.Errorf("Unexpected Vdo Error: %s", frame.Error.Error())
				continue
			}
			app.Syslog.Info(frame.String())

			// Send frame to server
			frame_message := &cctv.FrameMessage{Frame: frame.Data}
			if err := stream.Send(frame_message); err != nil {
				app.Syslog.Errorf("error sending frame: %v", err)
				if err == io.EOF {
					break loop
				}
				continue
			}

			// Receive acknowledgment from server
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

	if err := stream.CloseSend(); err != nil {
		app.Syslog.Errorf("error closing stream: %v", err)
	}
}
