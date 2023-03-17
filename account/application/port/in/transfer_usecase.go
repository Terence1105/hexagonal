package in

type ITransferUseCase interface {
	TransferUseCase(tranCom *TransferCommand) error
}
