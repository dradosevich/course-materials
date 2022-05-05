# Black Hat Go
## End of the Semester project repository 
## Danny Radosevich

### Things to note:
* As you advised I turn in, here is code I have alrady done for the project
* The encoder encodes an image
* There are two iterations of the decoder
* Includes the txt file of text to encode
* Also has the pre/post encoding picture
* encoded picture's md5sum should be cfe1d97c28dcadd2b9fbee60afe49ae3

### encoder.py
* Encodes a picture, currently picture and message are static
* creates a random offset of 8bytes to begin message
* Message encased in "!!!" on either side, this was for a challenge lab needed a marker
* Takes the bytes from the original picture, before going through and stripping the least significant bit before replacing it with bits from the message
* currently it encodes "ragtime cowboy joe" into a picture
### decode*.py
* These two scripts decode the image
* decoder.py goes through 8bytes at a time to find the markers before outputting the characters to the terminal
* decodetwo.py generates a string from the LSB of every byte in the picture, then uses regex to pull the string, knowing it is between "!!!" 

