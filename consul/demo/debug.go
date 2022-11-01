type debugArgs struct {
	Cluster string
	Server  string
	Index   int
	Debug   string
}

const debugFlag = "debug"

var debugsArgsIns *debugArgs

func initDebug() {
	debugsArgsIns = &debugArgs{}
	argsLen := len(os.Args)
	if !(argsLen > 3 && os.Args[argsLen-1] == debugFlag) {
		klog.Debug("本地启动参数有误")
		return
	}
	Index, err := strconv.Atoi(os.Args[argsLen-2])
	if err != nil {
		klog.Panic(err)
		return
	}
	debugsArgsIns = &debugArgs{
		Cluster: os.Args[argsLen-4],
		Server:  os.Args[argsLen-3],
		Index:   Index,
		Debug:   os.Args[argsLen-1],
	}
	if IsLocalRun() {
		klog.ResetToDevelopment()
	}
}

func InitDebugForTest() {
	debugsArgsIns = &debugArgs{
		Cluster: "re_dev",
		Server:  "re_logic",
		Index:   5,
		Debug:   "debug",
	}
	if IsLocalRun() {
		klog.ResetToDevelopment()
	}
}

func IsLocalRun() bool {
	if debugsArgsIns.Debug == debugFlag {
		return true
	}
	return false
}

func GetLocalArgsIndex() int {
	return debugsArgsIns.Index
}

func GetLocalClusterArgs() string {
	return debugsArgsIns.Cluster
}

func GetLocalServerArgs() string {
	return debugsArgsIns.Server
}
