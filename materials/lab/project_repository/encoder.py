# Danny Radosevich
# Encode a message into a picture
# using the LSB
import random
import re

old_pic = open("apollo.jpg","rb")
encoded = open("encoded_apollo.jpg","wb")
message = open("ragtime.txt","r")
#clean up the message
message = message.read()
regex = re.compile('[^a-zA-Z \n]')
message = regex.sub("",message)
message = message.lower()
#print(message)
#now we know where we are going to start writing out, can't make it too easy 
offset = random.randint(0,100)
try:
    offest_bytes = old_pic.read(8*offset)#read the fist 8*x bytes 
    #now we have the offset bytes 
    encoded.write(offest_bytes)
    #now we need to write our message
    #first the leading !
    
    for x in range(3):
        print(x)
        leading_bytes = old_pic.read(8)
        #print(ord('!'))# is 8
        leading = format(ord('!'),'#010b')[2:]
        #print(leading)
        i = 0
        for byte in leading_bytes:
            #hex_byte = hex(byte)
            byte = format(byte,'#010b')[2:]
            #print(hex_byte,"\t",byte)
            #print(type(byte))
            byte = list(byte)
            #print(len(byte), type(byte))
            byte[7] = leading[i]
            i+=1
            byte =  "".join(byte)
            byte = int(byte,2).to_bytes(1,"little")
            #print(byte)
            encoded.write(byte)
    #now we have to do the message 
    for c in message:
        #first get the ascii encoding 
        char_to_encode = ord(c)
        char_to_encode = format(char_to_encode,'#010b')[2:]
        bytes_to_encode = old_pic.read(8)
        i = 0
        print(c,end='')
        for byte in bytes_to_encode:
            #print(i)
            byte = format(byte,'#010b')[2:]
            byte = list(byte)
            byte[7] = char_to_encode[i]
            i+=1
            byte =  "".join(byte)
            byte = int(byte,2).to_bytes(1,"little")
            #print(byte)
            encoded.write(byte)
    #print("-------")
    for x in range(3):
        end_bytes = old_pic.read(8)
        ending = format(ord('!'),'#010b')[2:]
        i = 0
        for byte in end_bytes:
            byte = format(byte,'#010b')[2:]
            byte = list(byte)
            byte[7] = ending[i]
            i+=1
            byte =  "".join(byte)
            byte = int(byte,2).to_bytes(1,"little")
            #print(byte)
            encoded.write(byte)
    
    #print("--------------")
    while (byte := old_pic.read(1)):
        encoded.write(byte)
        

finally:
    old_pic.close()
    encoded.close()