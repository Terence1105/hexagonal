package main

func main() {

	// iTransferPort := adapter_out.NewAccountLevelDBAdapter()
	// iTransferUseCase := service.NewTransferService(iTransferPort)
	// adapter_in := adapter_in.NewGinAdapter(iTransferUseCase)
	// adapter_in.Run()

	gin := InitializeGinAdapter()
	gin.Run()

	println("Hello, world!")
}
