#Danny Radosevich
import re
encoded_pic = open("encoded_apollo.jpg","rb")
message = ""
while read_bytes := encoded_pic.read(8):
    b_char = ""
    for byte in read_bytes:
        b_char += str(byte%2)
    try:
        b_char = int(b_char,2)
        b_char = b_char.to_bytes(1, "little")
        b_char = b_char.decode()
        message = message + b_char 
    except:
        message = message
reg = "".join(re.findall(r"!!![a-z \n]*!*",message))
reg = reg.replace("!!!","")
reg = reg.replace("\\n","\n")
print(reg)