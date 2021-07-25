package week_9

import (
    "bufio"
    "io"
    "net"
)

type Protocol struct {
    PackageLength   uint
    HeaderLength    uint16
    ProtocolVersion uint16
    Operation       uint
    SequenceId      uint
    Body            []byte
}

func Decode(conn net.Conn) (Protocol, error) {
    reader := bufio.NewReader(conn)
    packageLength, err := readUint(reader)
    if err != nil {
        return Protocol{}, err
    }

    headerLength, err := readUint16(reader)
    if err != nil {
        return Protocol{}, err
    }

    protocolVersion, err := readUint16(reader)
    if err != nil {
        return Protocol{}, err
    }
    operation, err := readUint(reader)
    if err != nil {
        return Protocol{}, err
    }
    sequenceId, err := readUint(reader)
    if err != nil {
        return Protocol{}, err
    }
    data, err := readBytes(reader, packageLength+uint(headerLength))
    return Protocol{
        PackageLength:   packageLength,
        HeaderLength:    headerLength,
        ProtocolVersion: protocolVersion,
        Operation:       operation,
        SequenceId:      sequenceId,
        Body:            data,
    }, nil
}

func readUint(reader io.Reader) (uint, error) {
    buf := make([]byte, 4)
    _, err := io.ReadFull(reader, buf)
    if err != nil {
        return 0, err
    }
    var result uint
    result = uint(buf[3])<<24 | uint(buf[2])<<16 | uint(buf[1])<<8 | uint(buf[0])
    return result, nil
}

func readUint16(reader io.Reader) (uint16, error) {
    buf := make([]byte, 2)
    _, err := io.ReadFull(reader, buf)
    if err != nil {
        return 0, err
    }
    var result uint16
    result = uint16(buf[1])<<8 | uint16(buf[0])
    return result, nil
}

func readBytes(reader io.Reader, length uint) ([]byte, error) {
    buf := make([]byte, length)
    _, err := io.ReadFull(reader, buf)
    return buf, err
}
