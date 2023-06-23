module github.com/rfjakob/gocryptfs/v2

replace github.com/rfjakob/gocryptfs/v2/internal/fortanix => ./internal/fortanix/

go 1.16

require (
	github.com/aperturerobotics/jacobsa-crypto v1.0.0
	github.com/fortanix/sdkms-client-go v0.2.6
	github.com/hanwen/go-fuse/v2 v2.3.0
	github.com/moby/sys/mountinfo v0.6.2
	github.com/pkg/xattr v0.4.3
	github.com/rfjakob/eme v1.1.2
	github.com/sabhiram/go-gitignore v0.0.0-20201211210132-54b8a0bf510f
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035
)
