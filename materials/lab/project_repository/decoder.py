# Danny Radosevich
# decoder

encoded_pic = open("encoded_apollo.jpg","rb")
prev_string = ""
found = 0
while read_bytes := encoded_pic.read(8):
    b_char = ""
    for byte in read_bytes:
        byte = format(byte,'#010b')[2:]
        #print(byte)
        b_char += byte[7]#assign the last bit to our string 
    if b_char == format(ord('!'),'#010b')[2:]:
        found +=1 
    else:
        found = 0
    if found == 3:
        #print("Found")
        while message_bytes:= encoded_pic.read(8):
            c = ""
            for byte in message_bytes:
                byte = format(byte,'#010b')[2:]
                c += byte[7]
            if c != format(ord('!'),'#010b')[2:]:
                c = int(c,2)
                c = c.to_bytes(1,"little")
                c = c.decode()
                print(c,end="")
            else:
                break
    if found >3:
        break