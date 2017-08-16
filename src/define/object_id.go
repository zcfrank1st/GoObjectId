package define

import (
"os"
"bytes"
"encoding/binary"
"fmt"
"time"
"net"
"sync/atomic"
)

var (
    counter int64 = 0

    machineBytes []byte
    processBytes []byte

    now int64
    old int64 = 0
)

func init () {
    machineBytes = machine()
    processBytes = pid()
}

func ObjectId () []byte {
    t := timestamp()
    if now != old {
        atomic.StoreInt64(&counter, 0)
        old = now
    }

    c := count()

    objectIdBytes := make([]byte, 12)
    copy(objectIdBytes[:4], t)
    copy(objectIdBytes[4:7], machineBytes)
    copy(objectIdBytes[7:9], processBytes)
    copy(objectIdBytes[9:12], c)

    return objectIdBytes
}

func timestamp () []byte {
    now = time.Now().Unix()
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, now)
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    timestampBytes := make([]byte, 4)
    copy(timestampBytes, buf.Bytes())
    return timestampBytes
}

func machine () []byte {
    interfaces, err := net.Interfaces()
    var addr string
    if err == nil {
        for _, i := range interfaces {
            if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
                // Don't use random as we have a real address
                addr = i.HardwareAddr.String()
                break
            }
        }
    }

    machineBytes := make([]byte, 3)
    copy(machineBytes, []byte(addr))
    return machineBytes
}

func pid () []byte {
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, int64(os.Getpid()))
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    pidBytes := make([]byte, 2)
    copy(pidBytes, buf.Bytes())
    return pidBytes
}

func count() []byte {
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, atomic.AddInt64(&counter, 1))
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    countBytes := make([]byte, 3)
    copy(countBytes, buf.Bytes())
    return countBytes
}