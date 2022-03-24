import re

apicalls = open("api")
apicalls = apicalls.read()
#now we need to pull the request blocks
#first we need to pattern match all the blocks starting with <div class="request">
#terminating with</div>+<div class

pattern = r"<div class=\"request\">([\S\n\t\v ]*)</div>"
request_blocks = re.findall(pattern, apicalls) #it will only find one match not the 45 cause regex is dumb
#print(len(request_blocks)) #sanity check
request_blocks = request_blocks[0].split("<div class=\"request\">") #break apart the sections
#print(len(request_blocks)) # see we have more sections
for i in range(len(request_blocks)):
    request_blocks[i] = request_blocks[i].replace("\\n","\n") #get actual new line characters so I can read it
#print(request_blocks[0]) # sanity check round two
for block in request_blocks:
    #print("-----------------------------------------------------------------------")
    #now I need to actually split my data out
    # I no get the following information out:
    # requestUrl End point
    # number of parameters
    # Parameter1Name
    # P1 Type
    # ... through p4
    # Example URL
    block = block.replace("\n</li>","</li>")
    #print(block)
    #request_url_pattern = r"<pre class=\"prettyprint\"><code class=\"language-http\">(.*)</code>"
    request_url_pattern = r"\">/(.*?)</span>"
    rue = re.findall(request_url_pattern,block) #incidentally rue is my feeling about starting this
    if len(rue) > 0:
        rue = rue[0]
    print("https://apt.shodan.io/"+rue+", ",end='') #sanity check the RUE 
    para_pat = r"(<ul class=\"parameters\">[\n\W\w]*?<\/ul>)"
    para = re.findall(para_pat,block)
    for p in para:
        param_pat = r"<li>[\n\W\w]*?<\/li>"
        params = re.findall(param_pat,p)
        if len(params)>0:
            #params = params[0]
            #params = params.replace("<li>","")
            #params = params.split("</li>")
            print(str(len(params))+", ",end='')
            for pp in params:
                #print("*",pp)
                parname = r"<strong>(.*?):"
                pn = re.findall(parname,pp)
                #print(pn)
                if len(pn) > 0:
                    pn = pn[0]
                    print(pn+", ",end='')
                partype = r"\[(.*?)\]"
                pt = re.findall(partype,pp)
                #print(pt)
                if len(pt)>0:
                    pt = pt[0]
                    print(pt+", ",end='')
            if len(params) < 4:
                for i in range(4-len(params)):
                    print("nil, nil, ",end='')
    if(len(para))==0:
        print("0, nil, nil, nil, nil, nil, nil, nil, nil,",end='')
    #now finall we need to get the example url 
    #print(block)
    ex_pat = r"curl -[A-Z] [A-Z]+ \"([\S\n\t\v ]*?)\<\/code>"
    #print(ex_pat)
    exurl = re.findall(ex_pat,block)
    #print("\n",exurl)
    if len(exurl)>0:
        exurl = exurl[0]
        exurl = exurl.replace("\n","")
        print(exurl,end='')
    #para = str(para)
    print()
    