package main

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"github.com/fullsailor/pkcs7"
	"github.com/sassoftware/relic/v7/lib/authenticode"
	"os"
)
// https://docs.microsoft.com/en-us/dotnet/framework/tools/signtool-exe
//https://github.com/sassoftware/relic
//https://github.com/faradayio/vault-1
//https://github.com/Rillke/cert-chain-resolver
//https://github.com/jdferrell3/peinfo-go
 //https://github.com/Rillke/cert-chain-resolver/blob/master/certUtil/io.go
func main(){
	inParam := flag.String("in", "filename", "This specifies the input Signed PE filename to read from")
	//outParam := flag.String("out", "filename", "This specifies the output PKCS#7 filename to write to")

	flag.Parse()
	if *inParam == "filename" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	f,err:= os.Open(*inParam)
	if err != nil {
		panic(err)
	}

	r,err := authenticode.VerifyPE(f,true)
	if err != nil {
		panic(err)
	}
	//pkcs9.TimestampedSignature
	//	Indirect      *SpcIndirectDataContentPe
	//	ImageHashFunc crypto.Hash
	//	PageHashes    []byte
	//	PageHashFunc  crypto.Hash
	f1,_:= os.OpenFile("/Users/jianfenliu/Workspace/FileType/Xshell-7.0.0099p.exe", os.O_RDWR, 0644)
	err = authenticode.FixPEChecksum(f1)
	fmt.Printf("FixPEChecksumerr=%v\n",err)
	for _, re := range r{
		//fmt.Printf("Certificate=%v\n",re.Certificate)
		//fmt.Printf("re=%v\n",re.SignerInfo)
		//fmt.Printf("re=%v\n",re.Signature.VerifyChain)
		fmt.Printf("Subject=%v\n",re.Certificate.Subject)
		fmt.Printf("Issuer=%v\n",re.Certificate.Issuer)
		fmt.Printf("ImageHashFunc=%v\n",re.ImageHashFunc)
		fmt.Printf("PageHashFunc=%v\n",re.PageHashFunc)
		fmt.Printf("DigestAlgorithm.Algorithm=%v\n",re.SignerInfo.DigestAlgorithm.Algorithm)
		fmt.Printf("DigestAlgorithm.Parameters=%v\n",re.SignerInfo.DigestAlgorithm.Parameters)
		//fmt.Printf("PublicKey=%v\n",re.Certificate.PublicKey)
		fmt.Printf("PublicKeyAlgorithm=%v\n",re.Certificate.PublicKeyAlgorithm)
		fmt.Printf("PublicKeyAlgorithm String=%v\n",re.Certificate.PublicKeyAlgorithm.String())
		fmt.Printf("CounterSignature Hash=%v\n",re.CounterSignature.Hash)
		//fmt.Printf("CounterSignature Signature=%v\n",re.CounterSignature.Certificate.Signature)
		aaa := sha1.Sum(re.CounterSignature.Certificate.Signature)
		bbb := hex.EncodeToString(aaa[0:])
		fmt.Printf("CounterSignature Signature=%v\n",bbb)
		80e0b9c0c2a57ad2689ab5194773f129d4285d6b
		//Specifies the SHA1 hash of the signing certificate. The SHA1 hash is commonly specified
		//when multiple certificates satisfy the criteria specified by the remaining switches
		// 多个签名证书的时候会指定
		sha1byte:=sha1.Sum(re.Certificate.Raw)
		sha1str:=hex.EncodeToString(sha1byte[0:])
		fmt.Printf("sha1=%v\n",sha1str)

	}
	//
	//buf, err := signtool.ExtractDigitalSignature(*inParam)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if *outParam != "filename" {
	//	ioutil.WriteFile(*outParam, buf, 0644)
	//} else {
	//	fileName := filepath.Base(*inParam)
	//	if fileName == "." {
	//		fmt.Println("Input file path is not correct.")
	//		os.Exit(1)
	//	}
	//
	//	ioutil.WriteFile(fileName+".pkcs7", buf, 0644)
	//	certInfo,err := DecodeCertificate(buf)
	//	if err != nil {
	//		fmt.Printf("DecodeCertificate err%v\n", err)
	//	}
	//	fmt.Printf("certInfo=%v",certInfo)
	//}
	//attrs, err := pecert.GetAttributeCertificatesFromPath(*inParam)
	//if err != nil {
	//	fmt.Printf("%s\n", err)
	//}
	//for i, attr := range attrs {
	//	fmt.Printf("%d: %v\n",i, attr)
	//}
	//pefile, err := pe.NewPEFile(*inParam)
	//if err != nil {
	//	log.Println("Ooopss looks like there was a problem")
	//	log.Println(err)
	//	return
	//}
	//
	//log.Println("Imphash : ", pefile.GetImpHash())
	//
	//for _, section := range pefile.Sections {
	//	fmt.Println("-------------------------")
	//	data := pefile.GetData(section)
	//
	//	name := fmt.Sprintf("%s", section.Data.Name)
	//	md5 := section.Get_hash_md5(data)
	//	sha256 := section.Get_hash_sha256(data)
	//	entropy := section.Get_entropy(data)
	//	fmt.Println("name:", name)
	//	fmt.Println("md5 : ", md5)
	//	fmt.Println("sha256:", sha256)
	//	fmt.Println("entropy:", entropy)
	//}
}
func DecodeCertificate(data []byte) (*x509.Certificate, error) {
	//if IsPEM(data) {
	//	block, _ := pem.Decode(data)
	//	if block == nil {
	//		return nil, errors.New("Invalid certificate.")
	//	}
	//	if block.Type != certBlockType {
	//		return nil, errors.New("Invalid certificate.")
	//	}
	//
	//	data = block.Bytes
	//}

	cert, err := x509.ParseCertificate(data)
	if err == nil {
		return cert, nil
	}
	p, err := pkcs7.Parse(data)
	if err == nil {
		return p.Certificates[0], nil
	}

	return nil, errors.New("Invalid certificate.")
}
// https://github.com/vimeo/go-magic/tree/master/magic look
//func main()  {
//	//mtype := mimetype.Detect([]byte('d'))
//	// OR
//	//mtype, err := mimetype.DetectReader(io.Reader)
//	// OR
//	//mtype, err := mimetype.DetectFile("/opt/homebrew/Cellar/libmagic/5.40/share/misc/magic_ b.mgc")
//	//zipFile:= "/Users/jianfenliu/Workspace/FileType/testdata/7z.7z"
//	//mimeType, _ := mimetype.DetectFile(zipFile)
//	//log.Printf("mtype.String()=%v\n",mimeType.String())
//
//	decoder, err := magicmime.NewDecoder(magicmime.MAGIC_ERROR| magicmime.MAGIC_MIME|magicmime.MAGIC_DEBUG)
//
//
//
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	filePath:= "/Users/jianfenliu/Workspace/FileType/testdata/"
//	fileinfo,err := ioutil.ReadDir(filePath)
//	if err != nil{
//		log.Fatal(err)
//	}
//	for _, info := range fileinfo {
//		//log.Println(info.Name())
//		//t,err := decoder.TypeByFile(filePath+info.Name())
//		//buffer,err := ioutil.ReadFile(filePath+info.Name())
//		//if err != nil {
//		//	log.Println(err.Error())
//		//}
//		//t1, err := decoder.TypeByBuffer(buffer)
//		if err != nil {
//			log.Println(err.Error())
//		}
//		//fmt.Println("====================")
//		//fmt.Printf("libmagic.GetDefaultDir()=%s\n",libmagic.GetDefaultDir())
//
//		//libmagic.Load(libmagic.Open(libmagic.MAGIC_NO_CHECK_BUILTIN), libmagic.GetDefaultDir())
//		//t2 := libmagic.MimeFromFile(info.Name()+info.Name())
//		fmt.Printf("file=%s\t,t2=%v\n",info.Name()+info.Name(), t2)
//		//fmt.Printf("file=%s\tbyFile=%s\tbyBuf=%s, t2=%v\n",info.Name()+info.Name(), t,t1, t2)
//
//
//
//		//fmt.Printf("file=%s\tbyFile=%s\tbyBuf=%s\n",info.Name()+info.Name(), t,t1)
//		break
//	}
//	fmt.Printf("version is: %d\n", magicmime.Version())
//
//
//}
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