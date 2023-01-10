package system

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/pkg/conv"
	"fgzs-single/pkg/util/cmdutil"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zeromicro/go-zero/core/logx"
	"runtime"
	"strconv"
	"time"
)

type StatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatLogic {
	return &StatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (l *StatLogic) Stat(req *types.StatReq) (resp *types.StatResp, err error) {
	resp = new(types.StatResp)
	npmVersion, _, err := cmdutil.ExecCommand("npm -v")
	if err != nil {
		return nil, err
	}
	nodeVersion, _, err := cmdutil.ExecCommand("node -v")
	if err != nil {
		return nil, err
	}
	resp.Runtime = types.Runtime{
		GoVersion:   runtime.Version(),
		NpmVersion:  npmVersion,
		NodeVersion: nodeVersion,
		Os:          runtime.GOOS,
		Arch:        runtime.GOARCH,
	}
	Cpu := types.Cpu{
		VendorID:  "",
		ModelName: "",
		Cores:     "",
		CoresLoad: nil,
	}
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	if len(cpuInfo) > 0 {
		Cpu.VendorID = cpuInfo[0].VendorID
		Cpu.ModelName = cpuInfo[0].ModelName
		Cpu.Cores = strconv.Itoa(int(cpuInfo[0].Cores))
	}
	coresLoad, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	Cpu.CoresLoad = coresLoad
	resp.Cpu = Cpu
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	resp.Memory = types.Memory{
		Total:       strconv.Itoa(int(memory.Total)/MB) + " MB",
		Used:        strconv.Itoa(int(memory.Used)/MB) + " MB",
		Available:   strconv.Itoa(int(memory.Available)/MB) + " MB",
		UsedPercent: strconv.Itoa(conv.Int(memory.UsedPercent)) + "%",
	}
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	resp.Disk = types.Disk{
		Total:       strconv.Itoa(int(usage.Total)/MB) + " MB",
		Used:        strconv.Itoa(int(usage.Used)/MB) + " MB",
		Available:   strconv.Itoa(int(usage.Free)/MB) + " MB",
		UsedPercent: strconv.Itoa(conv.Int(usage.UsedPercent)) + "%",
	}
	return
}
