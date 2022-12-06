package chain

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	boardingProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                "Tony",
		hasBoardingPass:     false,
		hasLuggage:          false,
		isPassIdentifyCheck: false,
		isPassSecurityCheck: false,
		isCompleteBoard:     false,
	}
	boardingProcessor.ProcessFor(passenger)
}

func BuildBoardingProcessorChain() BoardingProcessor {
	completeBoardNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completeBoardNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckNode := &luggageCheckInProcessor{}
	luggageCheckNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckNode)
	return boardingPassNode
}
