package protolog

import (
	"bytes"
	"encoding/json"
	"time"
	"unicode"

	"github.com/golang/protobuf/proto"
)

type textMarshaller struct {
	disableTime     bool
	disableNewlines bool
}

func newTextMarshaller(options ...TextMarshallerOption) *textMarshaller {
	textMarshaller := &textMarshaller{
		false,
		false,
	}
	for _, option := range options {
		option(textMarshaller)
	}
	return textMarshaller
}

func (t *textMarshaller) Marshal(entry *Entry) ([]byte, error) {
	return textMarshalEntry(
		entry,
		t.disableTime,
		t.disableNewlines,
	)
}

func textMarshalEntry(
	entry *Entry,
	disableTime bool,
	disableNewlines bool,
) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if entry.ID != "" {
		_, _ = buffer.WriteString(entry.ID)
		_ = buffer.WriteByte(' ')
	}
	if !disableTime {
		_, _ = buffer.WriteString(entry.Time.Format(time.RFC3339))
		_ = buffer.WriteByte(' ')
	}
	levelString := entry.Level.String()
	_, _ = buffer.WriteString(levelString)
	_ = buffer.WriteByte(' ')
	if len(levelString) == 4 {
		_ = buffer.WriteByte(' ')
	}
	// TODO(pedge): verify only one of Event, Message, WriterOutput?
	if entry.Event != nil {
		if err := textMarshalMessage(buffer, entry.Event); err != nil {
			return nil, err
		}
	}
	if entry.Message != "" {
		_, _ = buffer.WriteString(entry.Message)
	}
	if entry.WriterOutput != nil {
		_, _ = buffer.Write(trimRightSpaceBytes(entry.WriterOutput))
	}
	if len(entry.Contexts) > 0 {
		_ = buffer.WriteByte(' ')
		lenContexts := len(entry.Contexts)
		for i, context := range entry.Contexts {
			if err := textMarshalMessage(buffer, context); err != nil {
				return nil, err
			}
			if i != lenContexts-1 {
				_ = buffer.WriteByte(' ')
			}
		}
	}
	if len(entry.Fields) > 0 {
		_ = buffer.WriteByte(' ')
		data, err := json.Marshal(entry.Fields)
		if err != nil {
			return nil, err
		}
		_, _ = buffer.Write(data)
	}
	data := trimRightSpaceBytes(buffer.Bytes())
	if !disableNewlines {
		buffer = bytes.NewBuffer(data)
		_ = buffer.WriteByte('\n')
		return buffer.Bytes(), nil
	}
	return data, nil
}

func textMarshalMessage(buffer *bytes.Buffer, message proto.Message) error {
	if message == nil {
		return nil
	}
	_, _ = buffer.WriteString(messageName(message))
	_ = buffer.WriteByte(' ')
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = buffer.Write(data)
	return err
}

func trimRightSpaceBytes(b []byte) []byte {
	return bytes.TrimRightFunc(b, unicode.IsSpace)
}
