package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"

	"runtime"

	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var NATSSTREAM = "NATS"

type MessageStore struct {
	MSiduuid    string // message id
	MSalias     string // alias
	MShostname  string // hostname
	MSipadrs    string // ip address
	MSmacid     string // macids
	MSmessage   string // message payload
	MSnodeuuid  string // unique id
	MSdate      string // message date
	MSsubject   string // message subject
	MSos        string // device os
	MSsequence  uint64 // consumer sequence for secure delete
	MSelementid int    // order in array
}
type Natsjs struct {
	NatsConnect *nats.Conn
	Js          jetstream.Stream

	Jetstream jetstream.JetStream

	Obsmp3   nats.ObjectStore
	Obsmp4   nats.ObjectStore
	OnAirmp3 jetstream.KeyValue
	OnAirmp4 jetstream.KeyValue
	Ctx      context.Context
	Ctxcan   context.CancelFunc
}

type NatsjsREPORT struct {
	NatsConnectREPORT *nats.Conn
	JsREPORT          jetstream.Stream
	JetstreamREPORT   jetstream.JetStream
	CtxREPORT         context.Context
	CtxcanREPORT      context.CancelFunc
}

func NewNatsJS() error {
	nnjsd := new(Natsjs)
	nnjsctx, nnjsctxcan := context.WithTimeout(context.Background(), 2048*time.Hour)

	natsopts := nats.Options{
		//Name:           "OPTS-" + alias,
		Url:            NatsServer,
		Verbose:        true,
		TLSConfig:      docerts(),
		AllowReconnect: true,
		MaxReconnect:   -1,
		ReconnectWait:  2,
		PingInterval:   120 * time.Second,
		Timeout:        2400 * time.Hour,
		User:           NatsUser,
		Password:       NatsUserPassword,
	}
	natsconnect, natsconnecterr := natsopts.Connect()
	if natsconnecterr != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + natsconnecterr.Error())
		}
		log.Println("NewNatsJS  connect" + getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + natsconnecterr.Error())
	}

	njsjetstream, njsjetstreamerr := jetstream.New(natsconnect)
	if njsjetstreamerr != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-snd") + getLangsNats("ms-err7") + njsjetstreamerr.Error())
		}
		log.Println("NewNatsJS jetstreamnew ", getLangsNats("ms-eraj"), njsjetstreamerr)
	}

	msgjs, msgjserr := njsjetstream.Stream(nnjsctx, "MESSAGES")
	if msgjserr != nil {
		log.Println("NewNatsJS NATS MESSAGES", getLangsNats("ms-eraj"), msgjserr)

	}

	natsoptsobject := nats.Options{
		Name:           "SN-" + NatsAlias,
		Url:            NatsServer,
		Verbose:        true,
		TLSConfig:      docerts(),
		AllowReconnect: true,
		MaxReconnect:   -1,
		ReconnectWait:  2,
		PingInterval:   120 * time.Second,
		Timeout:        5 * time.Second,
		User:           NatsUser,
		Password:       NatsUserPassword,
	}
	natsconnectobject, natsconnecterrmissing := natsoptsobject.Connect()
	if natsconnecterrmissing != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + natsconnecterrmissing.Error())
		}
		log.Println("SetupNATS object connect" + getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + natsconnecterrmissing.Error())
	}

	jsctx, jsmissingerr := natsconnectobject.JetStream()
	if jsmissingerr != nil {
		log.Println("SetupNATS JetStream ", getLangsNats("ms-eraj"), jsmissingerr)

	}

	jsctx.CreateObjectStore(&nats.ObjectStoreConfig{
		Bucket:      "mp3",
		Description: "MP3Bucket",
		Storage:     nats.FileStorage,
	})

	mp3obs, mp3errobs := jsctx.ObjectStore("mp3")
	if mp3errobs != nil {
		log.Println("SetupNATS ObjectStore mp3 ", mp3errobs)
	}

	jsctx.CreateObjectStore(&nats.ObjectStoreConfig{
		Bucket:      "mp4",
		Description: "MP4Bucket",
		Storage:     nats.FileStorage,
	})

	mp4obs, mp4errobs := jsctx.ObjectStore("mp4")
	if mp4errobs != nil {
		log.Println("SetupNATS ObjectStore mp4 ", mp4errobs)
	}
	onairmp3, kveerr := njsjetstream.CreateOrUpdateKeyValue(nnjsctx, jetstream.KeyValueConfig{
		Bucket: "OnAirmp3",
	})
	if kveerr != nil {
		log.Println("SetupNATS MP3 err", kveerr)
	}

	onairmp4, kveerr := njsjetstream.CreateOrUpdateKeyValue(nnjsctx, jetstream.KeyValueConfig{
		Bucket: "OnAirmp4",
	})
	if kveerr != nil {
		log.Println("SetupNATS MP4 err", kveerr)
	}

	nnjsd.Ctxcan = nnjsctxcan
	nnjsd.Ctx = nnjsctx
	nnjsd.Jetstream = njsjetstream
	nnjsd.NatsConnect = natsconnect
	nnjsd.Js = msgjs
	nnjsd.OnAirmp3 = onairmp3
	nnjsd.OnAirmp4 = onairmp4

	nnjsd.Obsmp3 = mp3obs
	nnjsd.Obsmp4 = mp4obs
	NATS = nnjsd
	return nil
}

//var ms = MessageStore{}

// var devicefound = false
// var messageloopauth = true
var shortServerName string
var shortServerName1 string
var memoryStats runtime.MemStats
var NatsMessages = make(map[int]MessageStore)
var NatsMessagesIndex = make(map[string]bool)

// var NatsMessagesDevice = make(map[int]MessageStore)
var NatsMessagesIndexDevice = make(map[string]bool)

// var fyneFilterFound = false
var MessageToSend string
var myNatsLang = "eng"

var myLangsNats = map[string]string{
	"eng-fl-ll":       "NATS Language to Use eng or esp",
	"eng-ms-err2":     "NATS No Connection ",
	"spa-ms-err2":     "NATS sin Conexión ",
	"hin-ms-err2":     "NATS कोई कनेक्शन नहीं ",
	"eng-ms-carrier":  "Carrier",
	"spa-ms-carrier":  "Transportador",
	"वाहक-ms-carrier": "Carrier",
	"eng-ms-nhn":      "No Host Name ",
	"spa-ms-nhn":      "Sin Nombre de Host ",
	"hin-ms-nhn":      "कोई होस्ट नाम नहीं ",
	"eng-ms-hn":       "Host ",
	"spa-ms-hn":       "Nombre de Host ",
	"hin-ms-hn":       "मेज़बान ",
	"eng-ms-mi":       "Mac IDS",
	"spa-ms-mi":       "ID de Mac",
	"hin-ms-mi":       "मैक आईडीएस",
	"eng-ms-ad":       "Address",
	"spa-ms-ad":       "Direccion",
	"hin-ms-ad":       "पता",
	"eng-ms-ni":       "Node Id - ",
	"spa-ms-ni":       "ID de Nodo - ",
	"hin-ms-ni":       "नोड आईडी - ",
	"eng-ms-msg":      "Message Id - ",
	"spa-ms-msg":      "ID de Mensaje - ",
	"hin-ms-msg":      "संदेश आईडी - ",
	"eng-ms-on":       "On - ",
	"spa-ms-on":       "En - ",
	"hin-ms-on":       "पर - ",
	"eng-ms-err6-1":   "Recieved ",
	"spa-ms-err6-1":   "Recibida ",
	"hin-ms-err6-1":   "प्राप्त ",
	"eng-ms-err6-2":   " Messages ",
	"spa-ms-err6-2":   " Mensajes ",
	"hin-ms-err6-2":   " संदेशों ",
	"eng-ms-err6-3":   " Logs",
	"spa-ms-err6-3":   " Registros",
	"hin-ms-err6-3":   " लॉग्स",
	"eng-ms-err7":     " NATS Server Missing",
	"spa-ms-err7":     " Falta el servidor NATS",
	"hin-ms-err7":     " NATS सर्वर गायब है",
	"eng-ms-eraj":     "Erase JetStream ",
	"spa-ms-eraj":     "Borrar JetStream ",
	"hin-ms-eraj":     "जेटस्ट्रीम मिटाएँ ",

	"eng-ms-err8": " JSON Marshall",
	"spa-ms-err8": " Mariscal JSON",
	"hin-ms-err8": " JSON मार्शल",

	"eng-ms-con": "Connected",
	"spa-ms-con": "Conectada",
	"hin-ms-con": "जुड़े हुए",
	"eng-ms-dis": "Disconnected",
	"spa-ms-dis": "Desconectada",
	"hin-ms-dis": "डिस्कनेक्ट किया गया",

	"eng-ms-snd": "Send ",
	"spa-ms-snd": "Enviar ",
	"hin-ms-snd": "भेजना ",

	"eng-ms-mde": "Message Decode Error ",
	"spa-ms-mde": "Error de Decodificación de Mensaje ",
	"hin-ms-mde": "संदेश डिकोड त्रुटि ",

	"eng-ms-root": "nhnats.go docerts() rootCAs Error ",
	"spa-ms-root": "Error de CA Raíz de nhnats.go docerts() ",
	"hin-ms-root": "nhnats.go docerts() rootCAs त्रुटि ",

	"eng-ms-client": "nhnats.go docerts() client cert Error",
	"spa-ms-client": "Error de Certificado de Cliente de nhnats.go docerts()",
	"hin-ms-client": "nhnats.go docerts() क्लाइंट प्रमाणपत्र त्रुटि",

	"eng-ms-sece": "Security Erase ",
	"spa-ms-sece": "Borrado de Seguridad ",
	"hin-ms-sece": "सुरक्षा मिटाएँ ",
	"eng-ms-nnm":  "No New Messages On Server ",
	"spa-ms-nnm":  "No hay Mensajes Nuevos en el Servidor ",
	"hin-ms-nnm":  "सर्वर पर कोई नया संदेश नहीं ",
}

// return translation strings
var gnlvalue string
var gnlerr bool

func getLangsNats(mystring string) string {
	gnlvalue, gnlerr = myLangsNats[myNatsLang+"-"+mystring]
	if !gnlerr {
		return myNatsLang + "-" + mystring
	}
	return gnlvalue
}

func docerts() *tls.Config {
	RootCAs, _ := x509.SystemCertPool()
	if RootCAs == nil {
		RootCAs = x509.NewCertPool()
	}
	ok := RootCAs.AppendCertsFromPEM([]byte(NatsCaroot))
	if !ok {
		log.Println(getLangsNats("ms-root"))
	}
	Clientcert, err := tls.X509KeyPair([]byte(NatsClientcert), []byte(NatsClientkey))
	if err != nil {
		log.Println(getLangsNats("ms-client") + err.Error())
	}
	shortServerName = strings.ReplaceAll(NatsServer, "nats://", "")
	shortServerName1 = strings.ReplaceAll(shortServerName, ":4222", "")
	TLSConfig := &tls.Config{
		RootCAs:            RootCAs,
		Certificates:       []tls.Certificate{Clientcert},
		ServerName:         shortServerName1,
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
	}

	return TLSConfig
}

//var puterr error

func PutBucket(bucket string, id string, data []byte) error {
	if bucket == "mp3" {
		_, puterr := NATS.Obsmp3.PutBytes(id, data)
		//log.Println("Put Bucket putobj", putobj.Opts, "Uploaded", id, "to", bucket, "size", len(data))
		if puterr != nil {
			return puterr
		}
	}
	if bucket == "mp4" {
		_, puterr := NATS.Obsmp4.PutBytes(id, data)

		if puterr != nil {
			return puterr
		}
	}
	runtime.GC()
	runtime.ReadMemStats(&memoryStats)
	return nil
}

//var gbdata []byte
//var gberr error

func GetBucket(bucket, id, station string) []byte {

	if bucket == "mp3" {

		gbdata, gberr := NATS.Obsmp3.GetBytes(id)
		//gbdata, gberr := NATS.Obsmp3.GetBytes(id)

		if gberr != nil {
			log.Panicln("Bucket MP3 Missing " + " bucket " + bucket + " id " + id + " error: " + gberr.Error())
			Send("messages."+station, "Bucket MP3 Missing "+" bucket "+bucket+" id "+id+" error: "+gberr.Error(), "nats")
		}
		runtime.GC()
		runtime.ReadMemStats(&memoryStats)
		return gbdata
	}
	if bucket == "mp4" {
		gbdata, gberr := NATS.Obsmp4.GetBytes(id)

		if gberr != nil {
			Send("messages."+station, "Bucket MP4 Missing "+" bucket "+bucket+" id "+id+" errror: "+gberr.Error(), "nats")
		}
		runtime.GC()
		return gbdata
	}
	return nil
}
func TestBucket(bucket, id string) bool {

	if bucket == "mp3" {
		_, gberr := NATS.Obsmp3.GetBytes(id)

		if gberr == nil {
			return true
		}
	}
	if bucket == "mp4" {
		_, gberr := NATS.Obsmp4.GetBytes(id)

		if gberr == nil {
			return true
		}

	}
	return false
}

//var gbs *nats.ObjectInfo
//var gbserr error

func GetBucketSize(bucket, id string) uint64 {
	if id == "" || id == "INTRO" || id == "OUTRO" {
		return 0
	}

	if bucket == "mp3" {
		gbs, gbserr := NATS.Obsmp3.GetInfo(id)
		if gbserr == nil {
			return gbs.Size
		}
	}
	if bucket == "mp4" {
		gbs, gbserr := NATS.Obsmp4.GetInfo(id)
		if gbserr == nil {
			return gbs.Size
		}
	}

	return 0
}

func DeleteBucket(bucket, id string) error {
	if bucket == "mp3" {
		dbkverr := NATS.Obsmp3.Delete(id)
		if dbkverr != nil {
			Send("messages."+NatsAlias, "Delete Bucket mp3 "+bucket+"-"+id+" "+dbkverr.Error(), NatsAlias)
			return dbkverr
		}
	}
	if bucket == "mp4" {
		dbkverr := NATS.Obsmp4.Delete(id)
		if dbkverr != nil {
			Send("messages."+NatsAlias, "Delete Bucket mp4 "+bucket+"-"+id+" "+dbkverr.Error(), NatsAlias)

			return dbkverr
		}
	}

	return nil
}

// send message to nats
var sndctx context.Context
var sndctxcan context.CancelFunc

func Send(subject, m, alias string) bool {
	sndctx, sndctxcan = context.WithTimeout(context.Background(), 1*time.Minute)
	EncMessage := MessageStore{}
	EncMessage.MSsubject = subject
	EncMessage.MSos = runtime.GOOS
	name, err := os.Hostname()
	if err != nil {
		EncMessage.MShostname = getLangsNats("ms-nhn")
	} else {
		EncMessage.MShostname = name
	}

	ifas, err := net.Interfaces()
	if err != nil {
		EncMessage.MShostname += "-  " + getLangsNats("ms-carrier")
	}
	if err == nil {
		for _, ifa := range ifas {
			a := ifa.HardwareAddr.String()
			if a != "" {
				EncMessage.MSmacid += a + ", "
			}
		}

		addrs, _ := net.InterfaceAddrs()
		for _, addr := range addrs {
			EncMessage.MSipadrs += addr.String() + ", "
		}
	}
	EncMessage.MSalias = alias
	EncMessage.MSnodeuuid = NatsNodeUUID
	msiduuid := uuid.New().String()
	EncMessage.MSiduuid = msiduuid
	EncMessage.MSdate = time.Now().Format(time.UnixDate)
	EncMessage.MSmessage = m
	jsonmsg, jsonerr := json.Marshal(EncMessage)

	if jsonerr != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-err8") + jsonerr.Error())
		}
	}

	_, errpub := NATS.Jetstream.Publish(sndctx, subject, []byte(Encrypt(string(jsonmsg), NatsQueuePassword)))
	if errpub != nil {
		log.Println("nats.Send error", errpub)
	}
	sndctxcan()
	runtime.GC()
	return false
}

func SendONAIRmp3(m string) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, puterr := NATS.OnAirmp3.Put(ctx, "OnAirmp3", []byte(m))

	if puterr != nil {
		Send("messages."+NatsAlias, "Send On Air mp3 "+m+" "+puterr.Error(), NatsAlias)
	}
	cancel()

}
func SendONAIRmp4(m string) {
	//"station.mp3.*"
	log.Println(m)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, puterr := NATS.OnAirmp4.Put(ctx, "OnAirmp4", []byte(m))

	if puterr != nil {
		Send("messages."+NatsAlias, "Send On Air mp4 "+m+" "+puterr.Error(), NatsAlias)
	}
	cancel()

}

var natsoptsmissing nats.Options
var natsconnectmissing *nats.Conn
var connecterrmissing error
var jsmissingctx nats.JetStreamContext
var jsmissingerr error
var streammissing error
var mistrafficerr error

func SetupNATS() {
	// queue = NATS
	// subjects = TYPE.alias
	// TYPE = devices, events, auth, messages
	natsoptsmissing = nats.Options{
		Name:           "SN-" + NatsAlias,
		Url:            NatsServer,
		Verbose:        true,
		TLSConfig:      docerts(),
		AllowReconnect: true,
		MaxReconnect:   -1,
		ReconnectWait:  2,
		PingInterval:   3 * time.Second,
		Timeout:        5 * time.Second,
		User:           NatsUser,
		Password:       NatsUserPassword,
	}
	natsconnectmissing, connecterrmissing = natsoptsmissing.Connect()
	if connecterrmissing != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + connecterrmissing.Error())
		}
		log.Println("SetupNATS connect" + getLangsNats("ms-snd") + " " + getLangsNats("ms-err7") + connecterrmissing.Error())
	}

	jsmissingctx, jsmissingerr = natsconnectmissing.JetStream()
	if jsmissingerr != nil {
		log.Println("SetupNATS JetStream ", getLangsNats("ms-eraj"), jsmissingerr)

	}
	_, streammissing = jsmissingctx.StreamInfo("MESSAGES")
	if streammissing != nil {
		_, createerr := jsmissingctx.AddStream(&nats.StreamConfig{
			Name:      "MESSAGES",
			Subjects:  []string{"messages.*", "events.*", "authorizations.*", "devices.*"},
			Storage:   nats.FileStorage,
			MaxAge:    1024 * time.Hour,
			FirstSeq:  1,
			Retention: nats.LimitsPolicy,
		})
		if createerr != nil {
			log.Println("SetupNATS MESSAGES stream missing ", getLangsNats("ms-eraj"), streammissing)
		}
		_, mistrafficerr = jsmissingctx.AddStream(&nats.StreamConfig{
			Name:      "TRAFFIC",
			Subjects:  []string{"spins.*", "clicks.*", "onair.*"},
			Storage:   nats.FileStorage,
			MaxAge:    204800 * time.Hour,
			FirstSeq:  1,
			Retention: nats.LimitsPolicy,
		})
		if mistrafficerr != nil {
			log.Println("SetupNATS TRAFFIC stream missing ", getLangsNats("ms-eraj"), mistrafficerr)
		}

		_, misaudioerr := jsmissingctx.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket:      "mp3",
			Description: "MP3Bucket",
			Storage:     nats.FileStorage,
		})
		if misaudioerr != nil {
			log.Println("SetupNATS Audio Bucket ", misaudioerr)
		}

		_, misvideoerr := jsmissingctx.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket:      "mp4",
			Description: "MP4Bucket",
			Storage:     nats.FileStorage,
		})
		if misvideoerr != nil {
			log.Println("SetupNATS Video Bucket ", misvideoerr)
		}

	}

}

func ReceiveMESSAGE() {

	var startseqmsg uint64 = 1
	rmctx, rmctxcan := context.WithTimeout(context.Background(), 4096*time.Hour)
	defer rmctxcan()
	rmconsumer, rmconserr := NATS.Js.CreateOrUpdateConsumer(rmctx, jetstream.ConsumerConfig{
		Name: "RM-" + NatsAlias,
		//Durable:           subject + alias,
		AckPolicy:         jetstream.AckExplicitPolicy,
		DeliverPolicy:     jetstream.DeliverByStartSequencePolicy,
		InactiveThreshold: 1024 * time.Hour,
		FilterSubject:     "messages.*",
		ReplayPolicy:      jetstream.ReplayInstantPolicy,
		OptStartSeq:       startseqmsg,
	})
	if rmconserr != nil {
		log.Println("ReceiveMESSAGEadmin Consumer", rmconserr)
	}
	for {

		rmmsg, rmerr := rmconsumer.Next()

		startseqmsg++

		if rmerr == nil {
			rmmeta, _ := rmmsg.Metadata()
			runtime.GC()
			runtime.ReadMemStats(&memoryStats)
			rmmsg.Ack()
			ms := MessageStore{}
			jsaonerr := json.Unmarshal([]byte(string(Decrypt(string(rmmsg.Data()), NatsQueuePassword))), &ms)
			if jsaonerr != nil {
				// send decrypt
				if FyneMessageWin != nil {
					FyneMessageWin.SetTitle(getLangsNats("ms-mde"))
				}
				log.Println("ReceiveMESSAGEadmin Un Marhal", jsaonerr)
			}

			if !NatsMessagesIndex[ms.MSiduuid] {
				ms.MSsequence = rmmeta.Sequence.Stream
				ms.MSelementid = len(NatsMessages)
				NatsMessages[len(NatsMessages)] = ms
				NatsMessagesIndex[ms.MSiduuid] = true
				FyneMessageList.Refresh()
			}
			if FyneMainWin != nil {
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)
				FyneMessageWin.SetTitle(getLangsNats("ms-err6-1") + strconv.Itoa(len(NatsMessages)) + getLangsNats("ms-err6-2") + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
			}
			FyneMessageList.Refresh()
		}

	}
}

func ReceiveONAIRMP3() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()

	mp3msg, mp3err := NATS.OnAirmp3.Watch(ctx, "OnAirmp3")
	if mp3err != nil {
		log.Println("ReceiveONAIRMP3", mp3err)
		Send("messages."+NatsAlias, "Receive On Air mp3 ", NatsAlias)
	}
	for {

		kve := <-mp3msg.Updates()
		if kve != nil {
			runtime.GC()
			runtime.ReadMemStats(&memoryStats)
			if FyneMainWin != nil {
				FyneMainWin.SetTitle("On Air MP3 " + string(kve.Value()) + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
			}
		}
	}
}

func DeleteNatsMessage(seq uint64) {
	dnmctx, dnmctxcan := context.WithTimeout(context.Background(), 1*time.Minute)
	dnmerr := NATS.Js.SecureDeleteMsg(dnmctx, seq)

	if dnmerr != nil {
		if FyneMessageWin != nil {
			FyneMessageWin.SetTitle(getLangsNats("ms-snd") + getLangsNats("ms-err7") + dnmerr.Error())
		}
	}
	dnmctxcan()
}

func EraseMessages(queue string) {
	log.Println("nhnats.go Erase MessagesConnect", queue)
	nc, connecterr := nats.Connect(NatsServer, nats.UserInfo(NatsUser, NatsUserPassword), nats.Secure(docerts()))
	if connecterr != nil {
		log.Println("nhnats.go Erase Messages Connect", getLangsNats("ms-erac"), connecterr.Error())
	}
	js, jserr := nc.JetStream()
	if jserr != nil {
		log.Println("nhnats.go Erase Messages Jetstream Make ", getLangsNats("ms-eraj"), jserr)
	}

	jspurge := js.PurgeStream(queue)
	if jspurge != nil {
		log.Println("nhnats.go Erase Messages Jetstream Purge "+queue, getLangsNats("ms-dels"), jspurge)
	}

	nc.Close()
}
func SetupDetails(queue string, age string) {

	log.Println("nhnats.go Erase Connect", queue, " ", age)
	nc, connecterr := nats.Connect(NatsServer, nats.UserInfo(NatsUser, NatsUserPassword), nats.Secure(docerts()))
	if connecterr != nil {
		log.Println("nhnats.go Erase Connect", getLangsNats("ms-erac"), connecterr.Error())
	}
	js, jserr := nc.JetStream()
	if jserr != nil {
		log.Println("nhnats.go Erase Jetstream Make ", getLangsNats("ms-eraj"), jserr)
	}

	jspurge := js.PurgeStream(queue)
	if jspurge != nil {
		log.Println("nhnats.go Erase Jetstream Purge "+queue, getLangsNats("ms-dels"), jspurge)
	}
	jsdelete := js.DeleteStream(queue)
	if jsdelete != nil {
		log.Println("nhnats.go Erase Jetstream Delete "+queue, getLangsNats("ms-dels"), jsdelete)
	}

	msgmaxage, ageerr := time.ParseDuration("1024h")
	if ageerr != nil {
		log.Println("nhnats.go Erase Jetstream parse ", ageerr)
	}

	queuestr, queueerr := js.AddStream(&nats.StreamConfig{
		Name:     queue,
		Subjects: []string{strings.ToLower(queue)},
		Storage:  nats.FileStorage,
		MaxAge:   msgmaxage,
		FirstSeq: 1,
	})
	if queueerr != nil {
		log.Println("nhnats.go ", queue+" Addstream ", getLangsNats("ms-adds"), queueerr)
	}
	fmt.Printf(queue+": %v\n", queuestr)
	//Send(queue, strings.ToLower(queue), getLangsNats("ms-sece"), NatsAlias+":" +NatsNodeUUID+" created subject: " + queue)
	nc.Close()
}
