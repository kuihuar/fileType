package main

import (
	"FileType/v1/libmagic"
	"FileType/v1/magicmime"
	"fmt"
	//"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"log"
	//"os"

)
// https://github.com/vimeo/go-magic/tree/master/magic look
func main()  {
	//mtype := mimetype.Detect([]byte('d'))
	// OR
	//mtype, err := mimetype.DetectReader(io.Reader)
	// OR
	//mtype, err := mimetype.DetectFile("/opt/homebrew/Cellar/libmagic/5.40/share/misc/magic_ b.mgc")
	//zipFile:= "/Users/jianfenliu/Workspace/FileType/testdata/7z.7z"
	//mimeType, _ := mimetype.DetectFile(zipFile)
	//log.Printf("mtype.String()=%v\n",mimeType.String())

	//decoder, err := magicmime.NewDecoder(magicmime.MAGIC_ERROR| magicmime.MAGIC_MIME|magicmime.MAGIC_DEBUG)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	filePath:= "/Users/jianfenliu/Workspace/FileType/testdata/"
	fileinfo,err := ioutil.ReadDir(filePath)
	if err != nil{
		log.Fatal(err)
	}
	for _, info := range fileinfo {
		//log.Println(info.Name())
		//t,err := decoder.TypeByFile(filePath+info.Name())
		//buffer,err := ioutil.ReadFile(filePath+info.Name())
		//if err != nil {
		//	log.Println(err.Error())
		//}
		//t1, err := decoder.TypeByBuffer(buffer)
		//if err != nil {
		//	log.Println(err.Error())
		//}
		fmt.Println("====================")
		fmt.Printf("libmagic.GetDefaultDir()=%s\n",libmagic.GetDefaultDir())

		libmagic.Load(libmagic.Open(libmagic.MAGIC_NO_CHECK_BUILTIN), libmagic.GetDefaultDir())
		t2 := libmagic.MimeFromFile(info.Name()+info.Name())
		fmt.Printf("file=%s\t,t2=%v\n",info.Name()+info.Name(), t2)
		//fmt.Printf("file=%s\tbyFile=%s\tbyBuf=%s, t2=%v\n",info.Name()+info.Name(), t,t1, t2)



		//fmt.Printf("file=%s\tbyFile=%s\tbyBuf=%s\n",info.Name()+info.Name(), t,t1)
		break
	}
	fmt.Printf("version is: %d\n", magicmime.Version())


}
//fmt.Printf("mtype.Extension()=%v\n",mtype.Extension())

//mgc := libmagic.Open(libmagic.MAGIC_NONE)
//defer mgc.Close()
//mgc.SetFlags(libmagic.MAGIC_MIME | libmagic.MAGIC_MIME_ENCODING| libmagic.MAGIC_COMPRESS)
//mgcFile := libmagic.GetDefaultDir()
//if mgcFile == "" {
//	log.Println("err: gcFile is empty")
//}
//err := mgc.Check(mgcFile)
//if err != nil {
//	log.Fatalf("err: %s\n", err.Error())
//
//}
//err = mgc.Load(mgcFile)
//if err != nil {
//	log.Fatalf("err: %s\n", err.Error())
//}
////fmt.Printf("file: %s", mgc.File(os.Args[0]))
//mimeTypeString := mgc.File(zipFile)
//
//log.Printf("file: %s\n", mimeTypeString)
//
//log.Println(libmagic.Version())