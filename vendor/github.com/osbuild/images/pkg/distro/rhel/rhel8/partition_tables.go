package rhel8

import (
	"github.com/osbuild/images/internal/common"
	"github.com/osbuild/images/pkg/arch"
	"github.com/osbuild/images/pkg/disk"
	"github.com/osbuild/images/pkg/distro/rhel"
)

func defaultBasePartitionTables(t *rhel.ImageType) (disk.PartitionTable, bool) {
	switch t.Arch().Name() {
	case arch.ARCH_X86_64.String():
		return disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size:     1 * common.MebiByte,
					Bootable: true,
					Type:     disk.BIOSBootPartitionGUID,
					UUID:     disk.BIOSBootPartitionUUID,
				},
				{
					Size: 100 * common.MebiByte,
					Type: disk.EFISystemPartitionGUID,
					UUID: disk.EFISystemPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "vfat",
						UUID:         disk.EFIFilesystemUUID,
						Mountpoint:   "/boot/efi",
						FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
						FSTabFreq:    0,
						FSTabPassNo:  2,
					},
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Label:        "root",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}, true

	case arch.ARCH_AARCH64.String():
		return disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size: 100 * common.MebiByte,
					Type: disk.EFISystemPartitionGUID,
					UUID: disk.EFISystemPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "vfat",
						UUID:         disk.EFIFilesystemUUID,
						Mountpoint:   "/boot/efi",
						FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
						FSTabFreq:    0,
						FSTabPassNo:  2,
					},
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Label:        "root",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}, true

	case arch.ARCH_PPC64LE.String():
		return disk.PartitionTable{
			UUID: "0x14fc63d2",
			Type: "dos",
			Partitions: []disk.Partition{
				{
					Size:     4 * common.MebiByte,
					Type:     "41",
					Bootable: true,
				},
				{
					Size: 2 * common.GibiByte,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}, true

	case arch.ARCH_S390X.String():
		return disk.PartitionTable{
			UUID: "0x14fc63d2",
			Type: "dos",
			Partitions: []disk.Partition{
				{
					Size:     2 * common.GibiByte,
					Bootable: true,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}, true

	default:
		return disk.PartitionTable{}, false
	}
}

func edgeBasePartitionTables(t *rhel.ImageType) (disk.PartitionTable, bool) {
	switch t.Arch().Name() {
	case arch.ARCH_X86_64.String():
		return disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size:     1 * common.MebiByte,
					Bootable: true,
					Type:     disk.BIOSBootPartitionGUID,
					UUID:     disk.BIOSBootPartitionUUID,
				},
				{
					Size: 127 * common.MebiByte,
					Type: disk.EFISystemPartitionGUID,
					UUID: disk.EFISystemPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "vfat",
						UUID:         disk.EFIFilesystemUUID,
						Mountpoint:   "/boot/efi",
						Label:        "EFI-SYSTEM",
						FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
						FSTabFreq:    0,
						FSTabPassNo:  2,
					},
				},
				{
					Size: 384 * common.MebiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.FilesystemDataUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Mountpoint:   "/boot",
						Label:        "boot",
						FSTabOptions: "defaults",
						FSTabFreq:    1,
						FSTabPassNo:  1,
					},
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.LUKSContainer{
						Label:      "crypt_root",
						Cipher:     "cipher_null",
						Passphrase: "osbuild",
						PBKDF: disk.Argon2id{
							Memory:      32,
							Iterations:  4,
							Parallelism: 1,
						},
						Clevis: &disk.ClevisBind{
							Pin:              "null",
							Policy:           "{}",
							RemovePassphrase: true,
						},
						Payload: &disk.Filesystem{
							Type:         "xfs",
							Label:        "root",
							Mountpoint:   "/",
							FSTabOptions: "defaults",
							FSTabFreq:    0,
							FSTabPassNo:  0,
						},
					},
				},
			},
		}, true

	case arch.ARCH_AARCH64.String():
		return disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size: 127 * common.MebiByte,
					Type: disk.EFISystemPartitionGUID,
					UUID: disk.EFISystemPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "vfat",
						UUID:         disk.EFIFilesystemUUID,
						Mountpoint:   "/boot/efi",
						Label:        "EFI-SYSTEM",
						FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
						FSTabFreq:    0,
						FSTabPassNo:  2,
					},
				},
				{
					Size: 384 * common.MebiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.FilesystemDataUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Mountpoint:   "/boot",
						Label:        "boot",
						FSTabOptions: "defaults",
						FSTabFreq:    1,
						FSTabPassNo:  1,
					},
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.LUKSContainer{
						Label:      "crypt_root",
						Cipher:     "cipher_null",
						Passphrase: "osbuild",
						PBKDF: disk.Argon2id{
							Memory:      32,
							Iterations:  4,
							Parallelism: 1,
						},
						Clevis: &disk.ClevisBind{
							Pin:              "null",
							Policy:           "{}",
							RemovePassphrase: true,
						},
						Payload: &disk.Filesystem{
							Type:         "xfs",
							Label:        "root",
							Mountpoint:   "/",
							FSTabOptions: "defaults",
							FSTabFreq:    0,
							FSTabPassNo:  0,
						},
					},
				},
			},
		}, true

	default:
		return disk.PartitionTable{}, false
	}
}

func ec2PartitionTables(t *rhel.ImageType) (disk.PartitionTable, bool) {
	// x86_64 - without /boot
	// aarch  - <= 8.9 - 512MiB, 8.10 and centos: 1 GiB
	var aarch64BootSize uint64
	switch {
	case common.VersionLessThan(t.Arch().Distro().OsVersion(), "8.10") && t.IsRHEL():
		aarch64BootSize = 512 * common.MebiByte
	default:
		aarch64BootSize = 1 * common.GibiByte
	}

	x86PartitionTable := disk.PartitionTable{
		UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
		Type: "gpt",
		Partitions: []disk.Partition{
			{
				Size:     1 * common.MebiByte,
				Bootable: true,
				Type:     disk.BIOSBootPartitionGUID,
				UUID:     disk.BIOSBootPartitionUUID,
			},
			{
				Size: 200 * common.MebiByte,
				Type: disk.EFISystemPartitionGUID,
				UUID: disk.EFISystemPartitionUUID,
				Payload: &disk.Filesystem{
					Type:         "vfat",
					UUID:         disk.EFIFilesystemUUID,
					Mountpoint:   "/boot/efi",
					FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
					FSTabFreq:    0,
					FSTabPassNo:  2,
				},
			},
			{
				Size: 2 * common.GibiByte,
				Type: disk.FilesystemDataGUID,
				UUID: disk.RootPartitionUUID,
				Payload: &disk.Filesystem{
					Type:         "xfs",
					Label:        "root",
					Mountpoint:   "/",
					FSTabOptions: "defaults",
					FSTabFreq:    0,
					FSTabPassNo:  0,
				},
			},
		},
	}
	// RHEL EC2 x86_64 images prior to 8.9 support only BIOS boot
	if common.VersionLessThan(t.Arch().Distro().OsVersion(), "8.9") && t.IsRHEL() {
		x86PartitionTable = disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size:     1 * common.MebiByte,
					Bootable: true,
					Type:     disk.BIOSBootPartitionGUID,
					UUID:     disk.BIOSBootPartitionUUID,
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Label:        "root",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}
	}

	switch t.Arch().Name() {
	case arch.ARCH_X86_64.String():
		return x86PartitionTable, true

	case arch.ARCH_AARCH64.String():
		return disk.PartitionTable{
			UUID: "D209C89E-EA5E-4FBD-B161-B461CCE297E0",
			Type: "gpt",
			Partitions: []disk.Partition{
				{
					Size: 200 * common.MebiByte,
					Type: disk.EFISystemPartitionGUID,
					UUID: disk.EFISystemPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "vfat",
						UUID:         disk.EFIFilesystemUUID,
						Mountpoint:   "/boot/efi",
						FSTabOptions: "defaults,uid=0,gid=0,umask=077,shortname=winnt",
						FSTabFreq:    0,
						FSTabPassNo:  2,
					},
				},
				{
					Size: aarch64BootSize,
					Type: disk.FilesystemDataGUID,
					UUID: disk.FilesystemDataUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Mountpoint:   "/boot",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
				{
					Size: 2 * common.GibiByte,
					Type: disk.FilesystemDataGUID,
					UUID: disk.RootPartitionUUID,
					Payload: &disk.Filesystem{
						Type:         "xfs",
						Label:        "root",
						Mountpoint:   "/",
						FSTabOptions: "defaults",
						FSTabFreq:    0,
						FSTabPassNo:  0,
					},
				},
			},
		}, true

	default:
		return disk.PartitionTable{}, false
	}
}
