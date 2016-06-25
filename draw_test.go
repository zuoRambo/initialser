package initialser

import (
	"testing"
//"os"
//"log"
	"bufio"
	"os"
	"log"
)

func TestDraw(t *testing.T) {
	a := NewAvatar("a")
	a.FontSize = 80
	a.Font = "华文黑体"
	a.Size = 250
	a.Ext = "png"
	a.Background = "#666666"

	d, err := NewDrawer(a)
	if err != nil {
		t.Error("not expected error ,", err)
	}
	_, err = d.Draw()
	if err != nil {
		t.Error("not expected error ,", err)
	}
	_, err = d.DrawToBytes()
	if err != nil {
		t.Error("not expected error ,", err)
	}


	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	d.DrawToWriter(b)
	b.Flush();

}

