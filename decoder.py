import av 
import cv2


class H264Decoder():
    """
    A class to decode H.264 encoded video frames.
    
    Attributes:
        preview (bool): Flag to show the preview of decoded frames using OpenCV.
    """

    def __init__(self, preview=False):
        """
        Initializes the H264Decoder with the specified preview option.
        
        Args:
            preview (bool): If True, shows the preview of decoded frames.
        """
        # Create a codec context for H.264 decoding
        self.ctx = av.CodecContext.create('h264', 'r')
        # Flag to check if the context is initialized
        self.initalized = False
        # Store the preview flag
        self.preview = preview

    def show_frame(self, frame):
        """
        Displays a frame using OpenCV.
        
        Args:
            frame (av.VideoFrame): The video frame to be displayed.
        """
        # Convert the frame to a NumPy array with BGR format
        img = frame.to_ndarray(format='bgr24')
        # Display the image in a window
        cv2.imshow("GRPC H264 Preview", img)

    def decode(self, frame_message):
        """
        Decodes the given H.264 frame message.
        
        Args:
            frame_message: The frame message containing the video frame data.
        """
        # Initialize the codec context with the frame width and height if not already initialized
        if not self.initalized:
            self.ctx.width = frame_message.width
            self.ctx.height = frame_message.height
            self.initalized = True

        # Parse the frame message into packets
        packets = self.ctx.parse(frame_message.frame)

        # Iterate over each packet
        for pkt in packets:
            # Decode the packet into frames
            frames = self.ctx.decode(pkt)
            for frame in frames:
                # If preview is enabled, show the decoded frame
                if self.preview:
                    self.show_frame(frame)

            # If preview is enabled, handle the key press event to quit the preview
            if self.preview:
                if cv2.waitKey(1) & 0xFF == ord('q'):
                    break