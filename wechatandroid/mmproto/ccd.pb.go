// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ccd.proto

package mmproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ccd1 struct {
	StartTime            *uint64  `protobuf:"varint,1,req,name=startTime" json:"startTime,omitempty"`
	CheckTime            *uint64  `protobuf:"varint,2,req,name=checkTime" json:"checkTime,omitempty"`
	Count                *uint32  `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	EndTime              []uint64 `protobuf:"varint,4,rep,name=endTime" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ccd1) Reset()         { *m = Ccd1{} }
func (m *Ccd1) String() string { return proto.CompactTextString(m) }
func (*Ccd1) ProtoMessage()    {}
func (*Ccd1) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bf7797e9c0831d, []int{0}
}

func (m *Ccd1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ccd1.Unmarshal(m, b)
}
func (m *Ccd1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ccd1.Marshal(b, m, deterministic)
}
func (m *Ccd1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ccd1.Merge(m, src)
}
func (m *Ccd1) XXX_Size() int {
	return xxx_messageInfo_Ccd1.Size(m)
}
func (m *Ccd1) XXX_DiscardUnknown() {
	xxx_messageInfo_Ccd1.DiscardUnknown(m)
}

var xxx_messageInfo_Ccd1 proto.InternalMessageInfo

func (m *Ccd1) GetStartTime() uint64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *Ccd1) GetCheckTime() uint64 {
	if m != nil && m.CheckTime != nil {
		return *m.CheckTime
	}
	return 0
}

func (m *Ccd1) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *Ccd1) GetEndTime() []uint64 {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type Ccd2 struct {
	Checkid              *string  `protobuf:"bytes,1,req,name=checkid" json:"checkid,omitempty"`
	StartTime            *uint32  `protobuf:"varint,2,req,name=startTime" json:"startTime,omitempty"`
	CheckTime            *uint32  `protobuf:"varint,3,req,name=checkTime" json:"checkTime,omitempty"`
	Count1               *uint32  `protobuf:"varint,4,req,name=count1" json:"count1,omitempty"`
	Count2               *uint32  `protobuf:"varint,5,req,name=count2" json:"count2,omitempty"`
	Count3               *uint32  `protobuf:"varint,6,req,name=count3" json:"count3,omitempty"`
	Const1               *uint64  `protobuf:"varint,7,req,name=const1" json:"const1,omitempty"`
	Const2               *uint64  `protobuf:"varint,8,req,name=const2" json:"const2,omitempty"`
	Const3               *uint64  `protobuf:"varint,9,req,name=const3" json:"const3,omitempty"`
	Const4               *uint64  `protobuf:"varint,10,req,name=const4" json:"const4,omitempty"`
	Const5               *uint64  `protobuf:"varint,11,req,name=const5" json:"const5,omitempty"`
	Const6               *uint64  `protobuf:"varint,12,req,name=const6" json:"const6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ccd2) Reset()         { *m = Ccd2{} }
func (m *Ccd2) String() string { return proto.CompactTextString(m) }
func (*Ccd2) ProtoMessage()    {}
func (*Ccd2) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bf7797e9c0831d, []int{1}
}

func (m *Ccd2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ccd2.Unmarshal(m, b)
}
func (m *Ccd2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ccd2.Marshal(b, m, deterministic)
}
func (m *Ccd2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ccd2.Merge(m, src)
}
func (m *Ccd2) XXX_Size() int {
	return xxx_messageInfo_Ccd2.Size(m)
}
func (m *Ccd2) XXX_DiscardUnknown() {
	xxx_messageInfo_Ccd2.DiscardUnknown(m)
}

var xxx_messageInfo_Ccd2 proto.InternalMessageInfo

func (m *Ccd2) GetCheckid() string {
	if m != nil && m.Checkid != nil {
		return *m.Checkid
	}
	return ""
}

func (m *Ccd2) GetStartTime() uint32 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *Ccd2) GetCheckTime() uint32 {
	if m != nil && m.CheckTime != nil {
		return *m.CheckTime
	}
	return 0
}

func (m *Ccd2) GetCount1() uint32 {
	if m != nil && m.Count1 != nil {
		return *m.Count1
	}
	return 0
}

func (m *Ccd2) GetCount2() uint32 {
	if m != nil && m.Count2 != nil {
		return *m.Count2
	}
	return 0
}

func (m *Ccd2) GetCount3() uint32 {
	if m != nil && m.Count3 != nil {
		return *m.Count3
	}
	return 0
}

func (m *Ccd2) GetConst1() uint64 {
	if m != nil && m.Const1 != nil {
		return *m.Const1
	}
	return 0
}

func (m *Ccd2) GetConst2() uint64 {
	if m != nil && m.Const2 != nil {
		return *m.Const2
	}
	return 0
}

func (m *Ccd2) GetConst3() uint64 {
	if m != nil && m.Const3 != nil {
		return *m.Const3
	}
	return 0
}

func (m *Ccd2) GetConst4() uint64 {
	if m != nil && m.Const4 != nil {
		return *m.Const4
	}
	return 0
}

func (m *Ccd2) GetConst5() uint64 {
	if m != nil && m.Const5 != nil {
		return *m.Const5
	}
	return 0
}

func (m *Ccd2) GetConst6() uint64 {
	if m != nil && m.Const6 != nil {
		return *m.Const6
	}
	return 0
}

type Ccd3Body struct {
	Loc                  *uint32  `protobuf:"varint,1,req,name=loc" json:"loc,omitempty"`
	Root                 *uint32  `protobuf:"varint,2,req,name=root" json:"root,omitempty"`
	Debug                *uint32  `protobuf:"varint,3,req,name=debug" json:"debug,omitempty"`
	PackageSign          *string  `protobuf:"bytes,4,req,name=packageSign" json:"packageSign,omitempty"`
	RadioVersion         *string  `protobuf:"bytes,5,req,name=radioVersion" json:"radioVersion,omitempty"`
	BuildVersion         *string  `protobuf:"bytes,6,req,name=buildVersion" json:"buildVersion,omitempty"`
	DeviceId             *string  `protobuf:"bytes,7,req,name=deviceId" json:"deviceId,omitempty"`
	AndroidId            *string  `protobuf:"bytes,8,req,name=androidId" json:"androidId,omitempty"`
	SerialId             *string  `protobuf:"bytes,9,req,name=serialId" json:"serialId,omitempty"`
	Model                *string  `protobuf:"bytes,10,req,name=model" json:"model,omitempty"`
	CpuCount             *uint32  `protobuf:"varint,11,req,name=cpuCount" json:"cpuCount,omitempty"`
	CpuBrand             *string  `protobuf:"bytes,12,opt,name=cpuBrand" json:"cpuBrand,omitempty"`
	CpuExt               *string  `protobuf:"bytes,13,opt,name=cpuExt" json:"cpuExt,omitempty"`
	WlanAddress          *string  `protobuf:"bytes,14,opt,name=wlanAddress" json:"wlanAddress,omitempty"`
	Ssid                 *string  `protobuf:"bytes,15,opt,name=ssid" json:"ssid,omitempty"`
	Bssid                *string  `protobuf:"bytes,16,opt,name=bssid" json:"bssid,omitempty"`
	SimOperator          *string  `protobuf:"bytes,17,opt,name=simOperator" json:"simOperator,omitempty"`
	WifiName             *string  `protobuf:"bytes,18,opt,name=wifiName" json:"wifiName,omitempty"`
	BuildFP              *string  `protobuf:"bytes,19,opt,name=buildFP" json:"buildFP,omitempty"`
	BuildBoard           *string  `protobuf:"bytes,20,opt,name=buildBoard" json:"buildBoard,omitempty"`
	BuildBootLoader      *string  `protobuf:"bytes,21,opt,name=buildBootLoader" json:"buildBootLoader,omitempty"`
	BuildBrand           *string  `protobuf:"bytes,22,opt,name=buildBrand" json:"buildBrand,omitempty"`
	BuildDevice          *string  `protobuf:"bytes,23,opt,name=buildDevice" json:"buildDevice,omitempty"`
	BuildHardware        *string  `protobuf:"bytes,24,opt,name=buildHardware" json:"buildHardware,omitempty"`
	BuildProduct         *string  `protobuf:"bytes,25,opt,name=buildProduct" json:"buildProduct,omitempty"`
	BuildManufacturer    *string  `protobuf:"bytes,26,opt,name=buildManufacturer" json:"buildManufacturer,omitempty"`
	PhoneNum             *string  `protobuf:"bytes,27,opt,name=phoneNum" json:"phoneNum,omitempty"`
	NetType              *string  `protobuf:"bytes,28,opt,name=netType" json:"netType,omitempty"`
	Qemu                 *uint32  `protobuf:"varint,29,opt,name=qemu" json:"qemu,omitempty"`
	Modified             *uint32  `protobuf:"varint,30,opt,name=modified" json:"modified,omitempty"`
	Task                 *uint32  `protobuf:"varint,31,opt,name=task" json:"task,omitempty"`
	PackageName          *string  `protobuf:"bytes,32,opt,name=packageName" json:"packageName,omitempty"`
	AppName              *string  `protobuf:"bytes,33,opt,name=appName" json:"appName,omitempty"`
	DataDir              *string  `protobuf:"bytes,34,opt,name=dataDir" json:"dataDir,omitempty"`
	ClassLoader          *string  `protobuf:"bytes,35,opt,name=classLoader" json:"classLoader,omitempty"`
	HardwareMask         *uint32  `protobuf:"varint,38,opt,name=hardwareMask" json:"hardwareMask,omitempty"`
	Luckpackcount        *uint32  `protobuf:"varint,41,opt,name=luckpackcount" json:"luckpackcount,omitempty"`
	BaseAPKMD5           *string  `protobuf:"bytes,42,opt,name=baseAPKMD5" json:"baseAPKMD5,omitempty"`
	ClientVersion        *string  `protobuf:"bytes,43,opt,name=clientVersion" json:"clientVersion,omitempty"`
	TbVersion            *string  `protobuf:"bytes,44,opt,name=tbVersion" json:"tbVersion,omitempty"`
	Ip                   *string  `protobuf:"bytes,45,opt,name=ip" json:"ip,omitempty"`
	Locale               *string  `protobuf:"bytes,46,opt,name=locale" json:"locale,omitempty"`
	CallState            *uint32  `protobuf:"varint,47,opt,name=callState" json:"callState,omitempty"`
	KeyGuardSecure       *uint32  `protobuf:"varint,48,opt,name=keyGuardSecure" json:"keyGuardSecure,omitempty"`
	WifiOn               *uint32  `protobuf:"varint,50,opt,name=wifiOn" json:"wifiOn,omitempty"`
	XposeCall            *uint32  `protobuf:"varint,51,opt,name=xposeCall" json:"xposeCall,omitempty"`
	AdbEnable            *uint32  `protobuf:"varint,53,opt,name=adbEnable" json:"adbEnable,omitempty"`
	Monkey               *uint32  `protobuf:"varint,54,opt,name=monkey" json:"monkey,omitempty"`
	SplashName           *string  `protobuf:"bytes,55,opt,name=splashName" json:"splashName,omitempty"`
	OsBinderProxy        *string  `protobuf:"bytes,56,opt,name=osBinderProxy" json:"osBinderProxy,omitempty"`
	StubProxy            *string  `protobuf:"bytes,57,opt,name=stubProxy" json:"stubProxy,omitempty"`
	VirtualNet           *uint32  `protobuf:"varint,58,opt,name=virtualNet" json:"virtualNet,omitempty"`
	Vpn                  *uint32  `protobuf:"varint,59,opt,name=vpn" json:"vpn,omitempty"`
	SubScriberId         *string  `protobuf:"bytes,60,opt,name=subScriberId" json:"subScriberId,omitempty"`
	GsmSimSate           *string  `protobuf:"bytes,61,opt,name=gsmSimSate" json:"gsmSimSate,omitempty"`
	GsmSimOperator       *string  `protobuf:"bytes,62,opt,name=gsmSimOperator" json:"gsmSimOperator,omitempty"`
	GsmSimOperatorNumber *string  `protobuf:"bytes,63,opt,name=gsmSimOperatorNumber" json:"gsmSimOperatorNumber,omitempty"`
	SoterId              *string  `protobuf:"bytes,64,opt,name=soterId" json:"soterId,omitempty"`
	KernelReleaseNumber  *string  `protobuf:"bytes,65,opt,name=kernelReleaseNumber" json:"kernelReleaseNumber,omitempty"`
	UsbState             *uint32  `protobuf:"varint,66,opt,name=usbState" json:"usbState,omitempty"`
	Sign                 *string  `protobuf:"bytes,67,opt,name=sign" json:"sign,omitempty"`
	PackageFlag          *uint32  `protobuf:"varint,68,opt,name=packageFlag" json:"packageFlag,omitempty"`
	AccessFlag           *uint32  `protobuf:"varint,69,opt,name=accessFlag" json:"accessFlag,omitempty"`
	Unkonwn              *uint32  `protobuf:"varint,70,opt,name=unkonwn" json:"unkonwn,omitempty"`
	TbVersionCrc         *uint32  `protobuf:"varint,71,opt,name=tbVersionCrc" json:"tbVersionCrc,omitempty"`
	SfMD5                *string  `protobuf:"bytes,72,opt,name=sfMD5" json:"sfMD5,omitempty"`
	SfArmMD5             *string  `protobuf:"bytes,73,opt,name=sfArmMD5" json:"sfArmMD5,omitempty"`
	SfArm64MD5           *string  `protobuf:"bytes,74,opt,name=sfArm64MD5" json:"sfArm64MD5,omitempty"`
	SbMD5                *string  `protobuf:"bytes,75,opt,name=sbMD5" json:"sbMD5,omitempty"`
	SoterId2             *string  `protobuf:"bytes,76,opt,name=soterId2" json:"soterId2,omitempty"`
	WidevineDeviceID     *string  `protobuf:"bytes,77,opt,name=WidevineDeviceID" json:"WidevineDeviceID,omitempty"`
	FSID                 *string  `protobuf:"bytes,78,opt,name=FSID" json:"FSID,omitempty"`
	Oaid                 *string  `protobuf:"bytes,79,opt,name=oaid" json:"oaid,omitempty"`
	TimeCheck            *uint32  `protobuf:"varint,80,opt,name=timeCheck" json:"timeCheck,omitempty"`
	NanoTime             *uint32  `protobuf:"varint,81,opt,name=nanoTime" json:"nanoTime,omitempty"`
	Refreshtime          *uint32  `protobuf:"varint,83,opt,name=refreshtime" json:"refreshtime,omitempty"`
	SoftConfig           []byte   `protobuf:"bytes,84,opt,name=softConfig" json:"softConfig,omitempty"`
	SoftData             []byte   `protobuf:"bytes,85,opt,name=softData" json:"softData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ccd3Body) Reset()         { *m = Ccd3Body{} }
func (m *Ccd3Body) String() string { return proto.CompactTextString(m) }
func (*Ccd3Body) ProtoMessage()    {}
func (*Ccd3Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bf7797e9c0831d, []int{2}
}

func (m *Ccd3Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ccd3Body.Unmarshal(m, b)
}
func (m *Ccd3Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ccd3Body.Marshal(b, m, deterministic)
}
func (m *Ccd3Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ccd3Body.Merge(m, src)
}
func (m *Ccd3Body) XXX_Size() int {
	return xxx_messageInfo_Ccd3Body.Size(m)
}
func (m *Ccd3Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Ccd3Body.DiscardUnknown(m)
}

var xxx_messageInfo_Ccd3Body proto.InternalMessageInfo

func (m *Ccd3Body) GetLoc() uint32 {
	if m != nil && m.Loc != nil {
		return *m.Loc
	}
	return 0
}

func (m *Ccd3Body) GetRoot() uint32 {
	if m != nil && m.Root != nil {
		return *m.Root
	}
	return 0
}

func (m *Ccd3Body) GetDebug() uint32 {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return 0
}

func (m *Ccd3Body) GetPackageSign() string {
	if m != nil && m.PackageSign != nil {
		return *m.PackageSign
	}
	return ""
}

func (m *Ccd3Body) GetRadioVersion() string {
	if m != nil && m.RadioVersion != nil {
		return *m.RadioVersion
	}
	return ""
}

func (m *Ccd3Body) GetBuildVersion() string {
	if m != nil && m.BuildVersion != nil {
		return *m.BuildVersion
	}
	return ""
}

func (m *Ccd3Body) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Ccd3Body) GetAndroidId() string {
	if m != nil && m.AndroidId != nil {
		return *m.AndroidId
	}
	return ""
}

func (m *Ccd3Body) GetSerialId() string {
	if m != nil && m.SerialId != nil {
		return *m.SerialId
	}
	return ""
}

func (m *Ccd3Body) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Ccd3Body) GetCpuCount() uint32 {
	if m != nil && m.CpuCount != nil {
		return *m.CpuCount
	}
	return 0
}

func (m *Ccd3Body) GetCpuBrand() string {
	if m != nil && m.CpuBrand != nil {
		return *m.CpuBrand
	}
	return ""
}

func (m *Ccd3Body) GetCpuExt() string {
	if m != nil && m.CpuExt != nil {
		return *m.CpuExt
	}
	return ""
}

func (m *Ccd3Body) GetWlanAddress() string {
	if m != nil && m.WlanAddress != nil {
		return *m.WlanAddress
	}
	return ""
}

func (m *Ccd3Body) GetSsid() string {
	if m != nil && m.Ssid != nil {
		return *m.Ssid
	}
	return ""
}

func (m *Ccd3Body) GetBssid() string {
	if m != nil && m.Bssid != nil {
		return *m.Bssid
	}
	return ""
}

func (m *Ccd3Body) GetSimOperator() string {
	if m != nil && m.SimOperator != nil {
		return *m.SimOperator
	}
	return ""
}

func (m *Ccd3Body) GetWifiName() string {
	if m != nil && m.WifiName != nil {
		return *m.WifiName
	}
	return ""
}

func (m *Ccd3Body) GetBuildFP() string {
	if m != nil && m.BuildFP != nil {
		return *m.BuildFP
	}
	return ""
}

func (m *Ccd3Body) GetBuildBoard() string {
	if m != nil && m.BuildBoard != nil {
		return *m.BuildBoard
	}
	return ""
}

func (m *Ccd3Body) GetBuildBootLoader() string {
	if m != nil && m.BuildBootLoader != nil {
		return *m.BuildBootLoader
	}
	return ""
}

func (m *Ccd3Body) GetBuildBrand() string {
	if m != nil && m.BuildBrand != nil {
		return *m.BuildBrand
	}
	return ""
}

func (m *Ccd3Body) GetBuildDevice() string {
	if m != nil && m.BuildDevice != nil {
		return *m.BuildDevice
	}
	return ""
}

func (m *Ccd3Body) GetBuildHardware() string {
	if m != nil && m.BuildHardware != nil {
		return *m.BuildHardware
	}
	return ""
}

func (m *Ccd3Body) GetBuildProduct() string {
	if m != nil && m.BuildProduct != nil {
		return *m.BuildProduct
	}
	return ""
}

func (m *Ccd3Body) GetBuildManufacturer() string {
	if m != nil && m.BuildManufacturer != nil {
		return *m.BuildManufacturer
	}
	return ""
}

func (m *Ccd3Body) GetPhoneNum() string {
	if m != nil && m.PhoneNum != nil {
		return *m.PhoneNum
	}
	return ""
}

func (m *Ccd3Body) GetNetType() string {
	if m != nil && m.NetType != nil {
		return *m.NetType
	}
	return ""
}

func (m *Ccd3Body) GetQemu() uint32 {
	if m != nil && m.Qemu != nil {
		return *m.Qemu
	}
	return 0
}

func (m *Ccd3Body) GetModified() uint32 {
	if m != nil && m.Modified != nil {
		return *m.Modified
	}
	return 0
}

func (m *Ccd3Body) GetTask() uint32 {
	if m != nil && m.Task != nil {
		return *m.Task
	}
	return 0
}

func (m *Ccd3Body) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *Ccd3Body) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *Ccd3Body) GetDataDir() string {
	if m != nil && m.DataDir != nil {
		return *m.DataDir
	}
	return ""
}

func (m *Ccd3Body) GetClassLoader() string {
	if m != nil && m.ClassLoader != nil {
		return *m.ClassLoader
	}
	return ""
}

func (m *Ccd3Body) GetHardwareMask() uint32 {
	if m != nil && m.HardwareMask != nil {
		return *m.HardwareMask
	}
	return 0
}

func (m *Ccd3Body) GetLuckpackcount() uint32 {
	if m != nil && m.Luckpackcount != nil {
		return *m.Luckpackcount
	}
	return 0
}

func (m *Ccd3Body) GetBaseAPKMD5() string {
	if m != nil && m.BaseAPKMD5 != nil {
		return *m.BaseAPKMD5
	}
	return ""
}

func (m *Ccd3Body) GetClientVersion() string {
	if m != nil && m.ClientVersion != nil {
		return *m.ClientVersion
	}
	return ""
}

func (m *Ccd3Body) GetTbVersion() string {
	if m != nil && m.TbVersion != nil {
		return *m.TbVersion
	}
	return ""
}

func (m *Ccd3Body) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Ccd3Body) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *Ccd3Body) GetCallState() uint32 {
	if m != nil && m.CallState != nil {
		return *m.CallState
	}
	return 0
}

func (m *Ccd3Body) GetKeyGuardSecure() uint32 {
	if m != nil && m.KeyGuardSecure != nil {
		return *m.KeyGuardSecure
	}
	return 0
}

func (m *Ccd3Body) GetWifiOn() uint32 {
	if m != nil && m.WifiOn != nil {
		return *m.WifiOn
	}
	return 0
}

func (m *Ccd3Body) GetXposeCall() uint32 {
	if m != nil && m.XposeCall != nil {
		return *m.XposeCall
	}
	return 0
}

func (m *Ccd3Body) GetAdbEnable() uint32 {
	if m != nil && m.AdbEnable != nil {
		return *m.AdbEnable
	}
	return 0
}

func (m *Ccd3Body) GetMonkey() uint32 {
	if m != nil && m.Monkey != nil {
		return *m.Monkey
	}
	return 0
}

func (m *Ccd3Body) GetSplashName() string {
	if m != nil && m.SplashName != nil {
		return *m.SplashName
	}
	return ""
}

func (m *Ccd3Body) GetOsBinderProxy() string {
	if m != nil && m.OsBinderProxy != nil {
		return *m.OsBinderProxy
	}
	return ""
}

func (m *Ccd3Body) GetStubProxy() string {
	if m != nil && m.StubProxy != nil {
		return *m.StubProxy
	}
	return ""
}

func (m *Ccd3Body) GetVirtualNet() uint32 {
	if m != nil && m.VirtualNet != nil {
		return *m.VirtualNet
	}
	return 0
}

func (m *Ccd3Body) GetVpn() uint32 {
	if m != nil && m.Vpn != nil {
		return *m.Vpn
	}
	return 0
}

func (m *Ccd3Body) GetSubScriberId() string {
	if m != nil && m.SubScriberId != nil {
		return *m.SubScriberId
	}
	return ""
}

func (m *Ccd3Body) GetGsmSimSate() string {
	if m != nil && m.GsmSimSate != nil {
		return *m.GsmSimSate
	}
	return ""
}

func (m *Ccd3Body) GetGsmSimOperator() string {
	if m != nil && m.GsmSimOperator != nil {
		return *m.GsmSimOperator
	}
	return ""
}

func (m *Ccd3Body) GetGsmSimOperatorNumber() string {
	if m != nil && m.GsmSimOperatorNumber != nil {
		return *m.GsmSimOperatorNumber
	}
	return ""
}

func (m *Ccd3Body) GetSoterId() string {
	if m != nil && m.SoterId != nil {
		return *m.SoterId
	}
	return ""
}

func (m *Ccd3Body) GetKernelReleaseNumber() string {
	if m != nil && m.KernelReleaseNumber != nil {
		return *m.KernelReleaseNumber
	}
	return ""
}

func (m *Ccd3Body) GetUsbState() uint32 {
	if m != nil && m.UsbState != nil {
		return *m.UsbState
	}
	return 0
}

func (m *Ccd3Body) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func (m *Ccd3Body) GetPackageFlag() uint32 {
	if m != nil && m.PackageFlag != nil {
		return *m.PackageFlag
	}
	return 0
}

func (m *Ccd3Body) GetAccessFlag() uint32 {
	if m != nil && m.AccessFlag != nil {
		return *m.AccessFlag
	}
	return 0
}

func (m *Ccd3Body) GetUnkonwn() uint32 {
	if m != nil && m.Unkonwn != nil {
		return *m.Unkonwn
	}
	return 0
}

func (m *Ccd3Body) GetTbVersionCrc() uint32 {
	if m != nil && m.TbVersionCrc != nil {
		return *m.TbVersionCrc
	}
	return 0
}

func (m *Ccd3Body) GetSfMD5() string {
	if m != nil && m.SfMD5 != nil {
		return *m.SfMD5
	}
	return ""
}

func (m *Ccd3Body) GetSfArmMD5() string {
	if m != nil && m.SfArmMD5 != nil {
		return *m.SfArmMD5
	}
	return ""
}

func (m *Ccd3Body) GetSfArm64MD5() string {
	if m != nil && m.SfArm64MD5 != nil {
		return *m.SfArm64MD5
	}
	return ""
}

func (m *Ccd3Body) GetSbMD5() string {
	if m != nil && m.SbMD5 != nil {
		return *m.SbMD5
	}
	return ""
}

func (m *Ccd3Body) GetSoterId2() string {
	if m != nil && m.SoterId2 != nil {
		return *m.SoterId2
	}
	return ""
}

func (m *Ccd3Body) GetWidevineDeviceID() string {
	if m != nil && m.WidevineDeviceID != nil {
		return *m.WidevineDeviceID
	}
	return ""
}

func (m *Ccd3Body) GetFSID() string {
	if m != nil && m.FSID != nil {
		return *m.FSID
	}
	return ""
}

func (m *Ccd3Body) GetOaid() string {
	if m != nil && m.Oaid != nil {
		return *m.Oaid
	}
	return ""
}

func (m *Ccd3Body) GetTimeCheck() uint32 {
	if m != nil && m.TimeCheck != nil {
		return *m.TimeCheck
	}
	return 0
}

func (m *Ccd3Body) GetNanoTime() uint32 {
	if m != nil && m.NanoTime != nil {
		return *m.NanoTime
	}
	return 0
}

func (m *Ccd3Body) GetRefreshtime() uint32 {
	if m != nil && m.Refreshtime != nil {
		return *m.Refreshtime
	}
	return 0
}

func (m *Ccd3Body) GetSoftConfig() []byte {
	if m != nil {
		return m.SoftConfig
	}
	return nil
}

func (m *Ccd3Body) GetSoftData() []byte {
	if m != nil {
		return m.SoftData
	}
	return nil
}

type Ccd3 struct {
	Crc                  *uint32   `protobuf:"varint,1,req,name=crc" json:"crc,omitempty"`
	TimeStamp            *uint32   `protobuf:"varint,2,req,name=timeStamp" json:"timeStamp,omitempty"`
	Body                 *Ccd3Body `protobuf:"bytes,3,req,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ccd3) Reset()         { *m = Ccd3{} }
func (m *Ccd3) String() string { return proto.CompactTextString(m) }
func (*Ccd3) ProtoMessage()    {}
func (*Ccd3) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bf7797e9c0831d, []int{3}
}

func (m *Ccd3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ccd3.Unmarshal(m, b)
}
func (m *Ccd3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ccd3.Marshal(b, m, deterministic)
}
func (m *Ccd3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ccd3.Merge(m, src)
}
func (m *Ccd3) XXX_Size() int {
	return xxx_messageInfo_Ccd3.Size(m)
}
func (m *Ccd3) XXX_DiscardUnknown() {
	xxx_messageInfo_Ccd3.DiscardUnknown(m)
}

var xxx_messageInfo_Ccd3 proto.InternalMessageInfo

func (m *Ccd3) GetCrc() uint32 {
	if m != nil && m.Crc != nil {
		return *m.Crc
	}
	return 0
}

func (m *Ccd3) GetTimeStamp() uint32 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *Ccd3) GetBody() *Ccd3Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Ccd1)(nil), "mmproto.Ccd1")
	proto.RegisterType((*Ccd2)(nil), "mmproto.Ccd2")
	proto.RegisterType((*Ccd3Body)(nil), "mmproto.Ccd3Body")
	proto.RegisterType((*Ccd3)(nil), "mmproto.Ccd3")
}

func init() {
	proto.RegisterFile("ccd.proto", fileDescriptor_a5bf7797e9c0831d)
}

var fileDescriptor_a5bf7797e9c0831d = []byte{
	// 1258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x96, 0xdb, 0x77, 0xd4, 0xb6,
	0x13, 0xc7, 0x4f, 0x96, 0xe5, 0xb2, 0x4e, 0xc2, 0xc5, 0xf0, 0xe3, 0x37, 0xa5, 0x94, 0x6e, 0xd3,
	0x96, 0x93, 0x52, 0x9a, 0xc2, 0x86, 0xa4, 0xf7, 0x4b, 0xb2, 0x4b, 0x60, 0x0b, 0x09, 0xa9, 0x97,
	0xb6, 0x4f, 0x7d, 0x90, 0x25, 0x6d, 0xa2, 0xb3, 0xb6, 0xe4, 0x4a, 0x32, 0x61, 0xff, 0xab, 0x3e,
	0xf4, 0x0f, 0xec, 0x99, 0x91, 0xed, 0xf5, 0x02, 0x7d, 0xd3, 0xf7, 0x33, 0xe3, 0xb1, 0xfc, 0x95,
	0x46, 0x56, 0xd4, 0xe3, 0x5c, 0x6c, 0x15, 0xd6, 0x78, 0x13, 0x5f, 0xcc, 0x73, 0x1a, 0x6c, 0xd8,
	0xa8, 0x3b, 0xe4, 0xe2, 0x61, 0x7c, 0x3b, 0xea, 0x39, 0xcf, 0xac, 0x7f, 0xa9, 0x72, 0x09, 0x2b,
	0xfd, 0xce, 0x66, 0x37, 0x59, 0x00, 0x8c, 0xf2, 0x53, 0xc9, 0x67, 0x14, 0xed, 0x84, 0x68, 0x03,
	0xe2, 0x1b, 0xd1, 0x79, 0x6e, 0x4a, 0xed, 0xe1, 0x5c, 0xbf, 0xb3, 0xb9, 0x9e, 0x04, 0x11, 0x43,
	0x74, 0x51, 0x6a, 0x41, 0x4f, 0x74, 0xfb, 0xe7, 0x36, 0xbb, 0x49, 0x2d, 0x37, 0xfe, 0xe9, 0xd0,
	0x4b, 0x07, 0x98, 0x42, 0x55, 0x94, 0xa0, 0x57, 0xf6, 0x92, 0x5a, 0x2e, 0x4f, 0xa7, 0x43, 0x65,
	0xff, 0x6b, 0x3a, 0xe1, 0xa5, 0xad, 0xe9, 0xdc, 0x8c, 0x2e, 0xd0, 0x0c, 0x1e, 0x42, 0x97, 0x42,
	0x95, 0x6a, 0xf8, 0x00, 0xce, 0xb7, 0xf8, 0xa0, 0xe1, 0xdb, 0x70, 0xa1, 0xc5, 0xb7, 0x03, 0xd7,
	0xce, 0x3f, 0x84, 0x8b, 0xf4, 0xc5, 0x95, 0x6a, 0xf8, 0x00, 0x2e, 0xb5, 0xf8, 0xa0, 0xe1, 0xdb,
	0xd0, 0x6b, 0xf1, 0x45, 0x9d, 0x47, 0x10, 0xb5, 0xf8, 0xa3, 0x86, 0xef, 0xc0, 0x6a, 0x8b, 0xef,
	0x34, 0x7c, 0x17, 0xd6, 0x5a, 0x7c, 0x77, 0xe3, 0xef, 0x38, 0xba, 0x34, 0xe4, 0x62, 0x7b, 0xdf,
	0x88, 0x79, 0x7c, 0x35, 0x3a, 0x97, 0x19, 0x4e, 0xb6, 0xad, 0x27, 0x38, 0x8c, 0xe3, 0xa8, 0x6b,
	0x8d, 0xf1, 0x95, 0x5b, 0x34, 0xc6, 0x95, 0x11, 0x32, 0x2d, 0x4f, 0xea, 0x95, 0x21, 0x11, 0xf7,
	0xa3, 0xd5, 0x82, 0xf1, 0x19, 0x3b, 0x91, 0x13, 0x75, 0xa2, 0xc9, 0xa5, 0x5e, 0xd2, 0x46, 0xf1,
	0x46, 0xb4, 0x66, 0x99, 0x50, 0xe6, 0x77, 0x69, 0x9d, 0x32, 0x9a, 0x0c, 0xeb, 0x25, 0x4b, 0x0c,
	0x73, 0xd2, 0x52, 0x65, 0xa2, 0xce, 0xb9, 0x10, 0x72, 0xda, 0x2c, 0xbe, 0x15, 0x5d, 0x12, 0xf2,
	0x95, 0xe2, 0x72, 0x2c, 0xc8, 0xc4, 0x5e, 0xd2, 0x68, 0x5c, 0x44, 0xa6, 0x85, 0x35, 0x4a, 0x8c,
	0x05, 0x39, 0xd9, 0x4b, 0x16, 0x00, 0x9f, 0x74, 0xd2, 0x2a, 0x96, 0x8d, 0x05, 0xd9, 0xd9, 0x4b,
	0x1a, 0x8d, 0x5f, 0x95, 0x1b, 0x21, 0x33, 0xf2, 0xb3, 0x97, 0x04, 0x81, 0x4f, 0xf0, 0xa2, 0x1c,
	0xd2, 0x46, 0x5c, 0xa5, 0xcf, 0x6d, 0x74, 0x15, 0xdb, 0xb7, 0x4c, 0x0b, 0x58, 0xeb, 0xaf, 0x60,
	0xb5, 0x5a, 0x93, 0xdd, 0x45, 0xf9, 0xf8, 0xb5, 0x87, 0x75, 0x8a, 0x54, 0x0a, 0x5d, 0x3a, 0xcb,
	0x98, 0xde, 0x13, 0xc2, 0x4a, 0xe7, 0xe0, 0x32, 0x05, 0xdb, 0x08, 0x1d, 0x77, 0x4e, 0x09, 0xb8,
	0x42, 0x21, 0x1a, 0xe3, 0xdc, 0x52, 0x82, 0x57, 0x09, 0x06, 0x81, 0xb5, 0x9c, 0xca, 0x5f, 0x14,
	0xd2, 0x32, 0x6f, 0x2c, 0x5c, 0x0b, 0xb5, 0x5a, 0x08, 0x67, 0x78, 0xa6, 0xa6, 0xea, 0x88, 0xe5,
	0x12, 0xe2, 0x30, 0xc3, 0x5a, 0x63, 0x9b, 0x90, 0xab, 0x07, 0xc7, 0x70, 0x9d, 0x42, 0xb5, 0x8c,
	0xef, 0x44, 0x11, 0x0d, 0xf7, 0x0d, 0xb3, 0x02, 0x6e, 0x50, 0xb0, 0x45, 0xe2, 0xcd, 0xe8, 0x4a,
	0xa5, 0x8c, 0x7f, 0x6e, 0x98, 0x90, 0x16, 0xfe, 0x47, 0x49, 0x6f, 0xe2, 0x45, 0x25, 0xf2, 0xe8,
	0x66, 0xbb, 0x12, 0xb9, 0xd4, 0x8f, 0x56, 0x49, 0x8d, 0x68, 0xf9, 0xe0, 0xff, 0xe1, 0x0b, 0x5a,
	0x28, 0xfe, 0x24, 0x5a, 0x27, 0xf9, 0x94, 0x59, 0x71, 0xc6, 0xac, 0x04, 0xa0, 0x9c, 0x65, 0xd8,
	0xec, 0x9a, 0x63, 0x6b, 0x44, 0xc9, 0x3d, 0xbc, 0x47, 0x49, 0x4b, 0x2c, 0xbe, 0x1f, 0x5d, 0x23,
	0x7d, 0xc8, 0x74, 0x39, 0x65, 0xdc, 0x97, 0x56, 0x5a, 0xb8, 0x45, 0x89, 0x6f, 0x07, 0xd0, 0xb9,
	0xe2, 0xd4, 0x68, 0x79, 0x54, 0xe6, 0xf0, 0x7e, 0x70, 0xae, 0xd6, 0xe8, 0x9c, 0x96, 0xfe, 0xe5,
	0xbc, 0x90, 0x70, 0x3b, 0x38, 0x57, 0x49, 0x5c, 0xbb, 0xbf, 0x64, 0x5e, 0xc2, 0x07, 0xfd, 0x15,
	0xec, 0x16, 0x1c, 0x63, 0xa5, 0xdc, 0x08, 0x35, 0x55, 0x52, 0xc0, 0x1d, 0xe2, 0x8d, 0xc6, 0x7c,
	0xcf, 0xdc, 0x0c, 0x3e, 0x0c, 0xf9, 0x38, 0x6e, 0xf5, 0x11, 0x2d, 0x5b, 0x3f, 0x78, 0xd2, 0x42,
	0xf8, 0x7e, 0x56, 0x14, 0x14, 0xfd, 0x28, 0xbc, 0xbf, 0x92, 0x18, 0x11, 0xcc, 0xb3, 0x91, 0xb2,
	0xb0, 0x11, 0x22, 0x95, 0xc4, 0xaa, 0x3c, 0x63, 0xce, 0x55, 0xeb, 0xf5, 0x71, 0xa8, 0xda, 0x42,
	0xe8, 0xe1, 0x69, 0xe5, 0xe7, 0x21, 0xce, 0xe9, 0x2e, 0xcd, 0x69, 0x89, 0xe1, 0x6a, 0x64, 0x25,
	0x9f, 0xe1, 0x64, 0xc2, 0xd9, 0xfc, 0x19, 0x25, 0x2d, 0x43, 0x5a, 0x75, 0xe6, 0xe4, 0xde, 0xf1,
	0xb3, 0xc3, 0xd1, 0x0e, 0xdc, 0xab, 0x56, 0xbd, 0x21, 0x58, 0x85, 0x67, 0x4a, 0x6a, 0x5f, 0x37,
	0xf9, 0xe7, 0x61, 0x4d, 0x97, 0x20, 0x76, 0xb2, 0x4f, 0xeb, 0x8c, 0xfb, 0x94, 0xb1, 0x00, 0xf1,
	0xe5, 0xa8, 0xa3, 0x0a, 0xf8, 0x82, 0x70, 0x47, 0x15, 0xd8, 0x6f, 0x99, 0xe1, 0x2c, 0x93, 0xb0,
	0x15, 0xfa, 0x2d, 0x28, 0x3a, 0xd4, 0x59, 0x96, 0x4d, 0x3c, 0xf3, 0x12, 0xbe, 0xa4, 0xd9, 0x2e,
	0x40, 0x7c, 0x37, 0xba, 0x3c, 0x93, 0xf3, 0x27, 0x25, 0xb3, 0x62, 0x22, 0x79, 0x69, 0x25, 0x3c,
	0xa0, 0x94, 0x37, 0x28, 0x56, 0xc7, 0xbe, 0x79, 0xa1, 0x61, 0x40, 0xf1, 0x4a, 0x61, 0xf5, 0xd7,
	0x85, 0x71, 0x72, 0xc8, 0xb2, 0x0c, 0xb6, 0x43, 0xf5, 0x06, 0xd0, 0x59, 0x24, 0xd2, 0xc7, 0x9a,
	0xa5, 0x99, 0x84, 0x9d, 0x10, 0x6d, 0x00, 0xd6, 0xcc, 0x8d, 0x9e, 0xc9, 0x39, 0xec, 0x86, 0x9a,
	0x41, 0xa1, 0x7b, 0xae, 0xc8, 0x98, 0x3b, 0xa5, 0x05, 0xfe, 0x2a, 0xb8, 0xb7, 0x20, 0xe8, 0x9e,
	0x71, 0xfb, 0x4a, 0x0b, 0x69, 0x8f, 0xad, 0x79, 0x3d, 0x87, 0xaf, 0x83, 0x7b, 0x4b, 0x30, 0xfc,
	0xea, 0xca, 0x34, 0x64, 0x7c, 0x13, 0xdc, 0x6b, 0x00, 0xbe, 0xe3, 0x95, 0xb2, 0xbe, 0x64, 0xd9,
	0x91, 0xf4, 0xf0, 0x2d, 0xbd, 0xbf, 0x45, 0xf0, 0x3f, 0xf0, 0xaa, 0xd0, 0xf0, 0x1d, 0x05, 0x70,
	0x88, 0xbb, 0xc3, 0x95, 0xe9, 0x84, 0x5b, 0x95, 0x4a, 0x3b, 0x16, 0xf0, 0x7d, 0xe8, 0xb0, 0x36,
	0xc3, 0xaa, 0x27, 0x2e, 0x9f, 0xa8, 0x7c, 0x82, 0x66, 0xff, 0x10, 0x66, 0xbe, 0x20, 0xe8, 0x76,
	0x50, 0xcd, 0x91, 0xf5, 0x23, 0xe5, 0xbc, 0x41, 0xe3, 0x41, 0x74, 0x63, 0x99, 0x1c, 0x95, 0x79,
	0x2a, 0x2d, 0xfc, 0x44, 0xd9, 0xef, 0x8c, 0xe1, 0xce, 0x77, 0xc6, 0xd3, 0xd4, 0x7e, 0x0e, 0x3b,
	0xbf, 0x92, 0xf1, 0x83, 0xe8, 0xfa, 0x4c, 0x5a, 0x2d, 0xb3, 0x44, 0x66, 0x92, 0x39, 0x59, 0x15,
	0xdb, 0xa3, 0xac, 0x77, 0x85, 0xb0, 0x63, 0x4b, 0x97, 0x86, 0x2d, 0xb3, 0x1f, 0x3a, 0xb6, 0xd6,
	0x74, 0x3a, 0xe3, 0xef, 0x6d, 0x58, 0x9d, 0xce, 0xf8, 0x5f, 0x5b, 0x74, 0xec, 0x41, 0xc6, 0x4e,
	0x60, 0x44, 0x8f, 0xb4, 0x11, 0x3a, 0xc3, 0x38, 0x97, 0xce, 0x51, 0xc2, 0xe3, 0xe0, 0xf7, 0x82,
	0xe0, 0xec, 0x4b, 0x3d, 0x33, 0xfa, 0x4c, 0xc3, 0x01, 0x05, 0x6b, 0x89, 0xbe, 0x37, 0x9b, 0x7e,
	0x68, 0x39, 0x3c, 0x09, 0x5d, 0xd9, 0x66, 0xf8, 0x77, 0x70, 0x53, 0x6c, 0xb5, 0xa7, 0xe1, 0xef,
	0x40, 0x82, 0xfe, 0x75, 0xd3, 0x3d, 0x9b, 0x63, 0x60, 0x1c, 0x4e, 0xb0, 0x5a, 0xd3, 0x1e, 0xc3,
	0xf1, 0xee, 0x23, 0x8c, 0xfe, 0x52, 0xed, 0xb1, 0x86, 0x50, 0xc5, 0x14, 0x43, 0xcf, 0xaa, 0x8a,
	0x69, 0x5d, 0x31, 0x98, 0x3a, 0x80, 0xe7, 0x55, 0xc5, 0x4a, 0xc7, 0xf7, 0xa2, 0xab, 0x7f, 0x28,
	0xfc, 0x0b, 0x6b, 0x19, 0x4e, 0xee, 0xf1, 0x08, 0x0e, 0x29, 0xe7, 0x2d, 0x8e, 0x1e, 0x1e, 0x4c,
	0xc6, 0x23, 0x38, 0x0a, 0x1e, 0xe2, 0x18, 0x99, 0x61, 0x4a, 0xc0, 0x8b, 0xc0, 0x70, 0x4c, 0x27,
	0x80, 0xca, 0xe5, 0x10, 0xef, 0x60, 0x70, 0x1c, 0xfa, 0xa7, 0x01, 0x38, 0x1b, 0xcd, 0xb4, 0xa1,
	0xdb, 0xda, 0xaf, 0x61, 0x95, 0x6a, 0x8d, 0x2b, 0x62, 0xe5, 0xd4, 0x4a, 0x77, 0x8a, 0xf9, 0x30,
	0x09, 0x2b, 0xd2, 0x42, 0xe4, 0x80, 0x99, 0xfa, 0xa1, 0xd1, 0x53, 0x75, 0x02, 0x2f, 0xfb, 0x2b,
	0x9b, 0x6b, 0x49, 0x8b, 0x84, 0x6f, 0x9d, 0xfa, 0x11, 0xf3, 0x0c, 0x7e, 0xa3, 0x68, 0xa3, 0x37,
	0xfe, 0xa4, 0x8b, 0xe6, 0x36, 0x76, 0x09, 0xb7, 0xcd, 0x6d, 0x89, 0x5b, 0x5e, 0xcf, 0x78, 0xe2,
	0x59, 0x5e, 0xd4, 0x17, 0xcc, 0x06, 0xc4, 0x9f, 0x46, 0xdd, 0xd4, 0x88, 0x39, 0x5d, 0x9b, 0x56,
	0x07, 0xd7, 0xb6, 0xaa, 0xdb, 0xf2, 0x56, 0x7d, 0xfd, 0x4a, 0x28, 0xfc, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0x21, 0xc8, 0x9b, 0x51, 0x0b, 0x00, 0x00,
}
