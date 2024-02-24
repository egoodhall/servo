package ipc

import (
	"errors"
	"io"
	"os/exec"
)

func NewClientConn(cmd *exec.Cmd) (io.ReadWriteCloser, error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	return NewConn(stdout, stdin), nil
}

func NewConn(rd io.ReadCloser, wr io.WriteCloser) io.ReadWriteCloser {
	return &readWriterConn{
		rd: rd,
		wr: wr,
	}
}

type readWriterConn struct {
	rd io.ReadCloser
	wr io.WriteCloser
}

func (rwc *readWriterConn) Read(p []byte) (int, error) {
	return rwc.rd.Read(p)
}

func (rwc *readWriterConn) Write(p []byte) (int, error) {
	return rwc.wr.Write(p)
}

func (rwc *readWriterConn) Close() error {
	return errors.Join(
		rwc.rd.Close(),
		rwc.wr.Close(),
	)
}
