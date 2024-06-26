package target

const TargetNameWorkerServer TargetName = "org.osbuild.worker.server"

type WorkerServerTargetOptions struct{}

func (WorkerServerTargetOptions) isTargetOptions() {}

func NewWorkerServerTarget() *Target {
	return newTarget(TargetNameWorkerServer, &WorkerServerTargetOptions{})
}

func NewWorkerServerTargetResult(artifact *OsbuildArtifact) *TargetResult {
	return newTargetResult(TargetNameWorkerServer, nil, artifact)
}
