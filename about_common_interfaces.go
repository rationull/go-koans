package go_koans

import "bytes"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		//in.WriteTo(out) // this seems better, but this koan is about io.Reader and io.Writer...

		data := make([]byte, 11)
		in.Read(data)
		out.Write(data)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		data := make([]byte, 5)
		in.Read(data)
		out.Write(data)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
