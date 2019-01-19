package okex

import (
	"fmt"
)

// ErrResponse represents response body error.
type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.Code, e.Message)
}

const (
	// ErrCodeBodyCannotBeBlank body cannot be blank
	ErrCodeBodyCannotBeBlank = 30020
	// ErrCodeJSONDataFormatError json data format error
	ErrCodeJSONDataFormatError = 30021
	// ErrCodeParameterCannotBeBlank parameters returned respectively
	ErrCodeParameterCannotBeBlank = 30023
	// ErrCodeParameterValueError parameters returned respectively
	ErrCodeParameterValueError = 30024
	// ErrCodeParameterCategoryError parameter category error
	ErrCodeParameterCategoryError = 30025
	// ErrCodeRequestedTooFrequent endpointlimit exceeded; requested too frequent; endpoint limit exceeded
	ErrCodeRequestedTooFrequent = 30026
	// ErrCodeLoginFailure operating orders of other users
	ErrCodeLoginFailure = 30027
	// ErrCodeUnauthorizedExecution unauthorized execution
	ErrCodeUnauthorizedExecution = 30028
	// ErrCodeAccountSuspended account suspended
	ErrCodeAccountSuspended = 30029
	// ErrCodeEndpointRequestFailed please try again; endpoint request failed. Please try again
	ErrCodeEndpointRequestFailed = 30030
	// ErrCodeTokenNotExist token requested does not exist
	ErrCodeTokenNotExist = 30031
	// ErrCodePairNotExist pair requested does not exist
	ErrCodePairNotExist = 30032
	// ErrCodeExchangeDomainNotExist the error returned when the exchange for the apikey validation is not filled
	ErrCodeExchangeDomainNotExist = 30033
	// ErrCodeExchangeIDNotExist the error returned when the exchange ID for the apikey validation is not filled
	ErrCodeExchangeIDNotExist = 30034
	// ErrCodeTradingIsNotSupportedInThisWebsite the error returned when the exchange is closed
	ErrCodeTradingIsNotSupportedInThisWebsite = 30035
	// ErrCodeNoRelevantData no relevant data when enquiring the endpoint
	ErrCodeNoRelevantData = 30036
	// ErrCodeEndpointIsOfflineOrUnavailable endpoint is offline or unavailable
	ErrCodeEndpointIsOfflineOrUnavailable = 30037
	// ErrCodeUserNotExist user does not exist
	ErrCodeUserNotExist = 30038

	// ErrCodeFuturesAccountSuspended when the futures account is suspended
	ErrCodeFuturesAccountSuspended = 32001
	// ErrCodeFuturesAccountNotExist when futures trading is not enabled for the account
	ErrCodeFuturesAccountNotExist = 32002
	// ErrCodeCanceling please wait; when the user execute other operations during order cancellation
	ErrCodeCanceling = 32003
	// ErrCodeNoUnfilledOrders when the user check the unfilled orders
	ErrCodeNoUnfilledOrders = 32004
	// ErrCodeMaxOrderQuantity when the user placed an order exceeding the quantity limit
	ErrCodeMaxOrderQuantity = 32005
	// ErrCodeOrderPriceOrTriggerPriceExceedsUSDOneMillion when the user placed and order with price or trigger price over USD 1 million
	ErrCodeOrderPriceOrTriggerPriceExceedsUSDOneMillion = 32006
	// ErrCodeLeverageLevelMustBeTheSameForOrdersOnTheSameSideOfTheContract when the user has open positions with 10x leverage and tryingto open a 20x leverage order
	ErrCodeLeverageLevelMustBeTheSameForOrdersOnTheSameSideOfTheContract = 32007
	// ErrCodeMaxPositionsToOpenCrossMargin when the order quantity is larger than the availability (cross margin)
	ErrCodeMaxPositionsToOpenCrossMargin = 32008
	// ErrCodeMaxPositionsToOpenFixedMargin when the order quantity is larger than the availability fixed margin)
	ErrCodeMaxPositionsToOpenFixedMargin = 32009
	// ErrCodeLeverageCannotBeChangedWithOpenPositions if an user holds a short position with 10x leverage, he/she will not be able to change the leverage to 20x
	ErrCodeLeverageCannotBeChangedWithOpenPositions = 32010
	// ErrCodeFuturesStatusError contract expired
	ErrCodeFuturesStatusError = 32011
	// ErrCodeFuturesOrderUpdateError updating the status of a canceled order
	ErrCodeFuturesOrderUpdateError = 32012
	// ErrCodeTokenTypeIsBlank token type is blank
	ErrCodeTokenTypeIsBlank = 32013
	// ErrCodeNumberOfContractsClosingIsLargerThanTheNumberOfContractsAvailable your number of contracts closing is larger than the number of contracts available
	ErrCodeNumberOfContractsClosingIsLargerThanTheNumberOfContractsAvailable = 32014
	// ErrCodeMarginRatioIsLowerThanHundredPercentBeforeOpeningPositions margin ratio is lower than 100% before opening positions
	ErrCodeMarginRatioIsLowerThanHundredPercentBeforeOpeningPositions = 32015
	// ErrCodeMarginRatioIsLowerThanHundredPercentAfterOpeningPosition margin ratio is lower than 100% after opening position
	ErrCodeMarginRatioIsLowerThanHundredPercentAfterOpeningPosition = 32016
	// ErrCodeNoBBO no BBO
	ErrCodeNoBBO = 32017
	// ErrCodeOrderQuantityIsLessThanOne please try again; the order quantity is less than 1, please try again
	ErrCodeOrderQuantityIsLessThanOne = 32018
	// ErrCodeOrderPriceDeviatesFromThePriceOfThePreviousMinuteByMoreThanThreePercent the order price deviates from the price of the previous minute by more than 3%
	ErrCodeOrderPriceDeviatesFromThePriceOfThePreviousMinuteByMoreThanThreePercent = 32019
	// ErrCodePriceIsNotInTheRangeOfThePriceLimit the price is not in the range of the price limit
	ErrCodePriceIsNotInTheRangeOfThePriceLimit = 32020
	// ErrCodeLeverageError leverage is not set as 10x or 20x
	ErrCodeLeverageError = 32021
	// ErrCodeFunctionIsNotSupportedInCountryOrRegionAccordingToTheRegulations futures trading not supported in certain regions
	ErrCodeFunctionIsNotSupportedInCountryOrRegionAccordingToTheRegulations = 32022
	// ErrCodeAccountHasOutstandingLoan this account has outstanding loan
	ErrCodeAccountHasOutstandingLoan = 32023
	// ErrCodeOrderCannotBePlacedDuringDelivery order placement failure during delivery
	ErrCodeOrderCannotBePlacedDuringDelivery = 32024
	// ErrCodeOrderCannotBePlacedDuringSettlement order placement failure during settlement
	ErrCodeOrderCannotBePlacedDuringSettlement = 32025
	// ErrCodeAccountIsRestrictedFromOpeningPositions restricted from opening position
	ErrCodeAccountIsRestrictedFromOpeningPositions = 32026
	// ErrCodeCancelledOver20Orders order cancellation limit reached
	ErrCodeCancelledOver20Orders = 32027
	// ErrCodeAccountIsSuspendedAndLiquidated account is suspended and liquidated
	ErrCodeAccountIsSuspendedAndLiquidated = 32028
	// ErrCodeOrderInfoNotExist order has been canceled already
	ErrCodeOrderInfoNotExist = 32029

	// ErrCodeMarginAccountForThisPairIsNotEnabledYet the service must be enabled before trading
	ErrCodeMarginAccountForThisPairIsNotEnabledYet = 33001
	// ErrCodeMarginAccountForThisPairIsSuspended margin account suspended
	ErrCodeMarginAccountForThisPairIsSuspended = 33002
	// ErrCodeNoLoanBalance insufficient balance for loan
	ErrCodeNoLoanBalance = 33003
	// ErrCodeLoanAmountCannotBeSmallerThanTheMinimumLimit minimum loan amount limit not reached
	ErrCodeLoanAmountCannotBeSmallerThanTheMinimumLimit = 33004
	// ErrCodeRepaymentAmountMustExceedZero invalid repayment amount
	ErrCodeRepaymentAmountMustExceedZero = 33005
	// ErrCodeLoanOrderNotFound loan order not found
	ErrCodeLoanOrderNotFound = 33006
	// ErrCodeStatusNotFound status unchanged
	ErrCodeStatusNotFound = 33007
	// ErrCodeLoanAmountCannotExceedTheMaximumLimit invalid loan amount
	ErrCodeLoanAmountCannotExceedTheMaximumLimit = 33008
	// ErrCodeUserIDIsBlank user ID not provided
	ErrCodeUserIDIsBlank = 33009
	// ErrCodeCannotCancelAnOrderDuringSessionTwoOfCallAuction order cancellation not allowed during call auction
	ErrCodeCannotCancelAnOrderDuringSessionTwoOfCallAuction = 33010
	// ErrCodeNoNewMarketData no market data
	ErrCodeNoNewMarketData = 33011
	// ErrCodeOrderCancellationFailed order cancellation failed
	ErrCodeOrderCancellationFailed = 33012
	// ErrCodeOrderPlacementFailed order placement failed
	ErrCodeOrderPlacementFailed = 33013
	// ErrCodeOrderNotExist order canceled already. Invalid order number
	ErrCodeOrderNotExist = 33014
	// ErrCodeExceededMaximumLimit exceeded maximum limit during multiple-order placement
	ErrCodeExceededMaximumLimit = 33015
	// ErrCodeMarginTradingIsNotOpenForThisToken insufficient balance for order placement
	ErrCodeMarginTradingIsNotOpenForThisToken = 33016
	// ErrCodeMarginTradingInsufficientBalance margin trading not supported for this pair
	ErrCodeMarginTradingInsufficientBalance = 33017
	// ErrCodeParameterMustBeSmallerThanOne invalid parameter for getting market data
	ErrCodeParameterMustBeSmallerThanOne = 33018
	// ErrCodeRequestNotSupported margin trading not supported for some exchanges
	ErrCodeRequestNotSupported = 33020
	// ErrCodeTokenAndThePairDoNotMatch incorrect token for the token pair during repayment
	ErrCodeTokenAndThePairDoNotMatch = 33021
	// ErrCodePairAndTheOrderDoNotMatch incorrect token for the order during repayment
	ErrCodePairAndTheOrderDoNotMatch = 33022
	// ErrCodeCanOnlyPlaceMarketOrdersDuringCallAuction you can only place market orders during call auction
	ErrCodeCanOnlyPlaceMarketOrdersDuringCallAuction = 33023
	// ErrCodeTradingAmountTooSmall invalid trading amount
	ErrCodeTradingAmountTooSmall = 33024
	// ErrCodeBaseTokenAmountIsBlank settings not completed during order placement
	ErrCodeBaseTokenAmountIsBlank = 33025
	// ErrCodeTransactionCompleted cancel limited when the transaction completed
	ErrCodeTransactionCompleted = 33026
	// ErrCodeOrderCancelledOrCancelling cancel limited when the order is cancelling or cancelled
	ErrCodeOrderCancelledOrCancelling = 33027
	// ErrCodeDecimalPlacesOfTheTradingPriceExceededTheLimit order endpoint: The decimal places of the trading price exceeded the limit
	ErrCodeDecimalPlacesOfTheTradingPriceExceededTheLimit = 33028
	// ErrCodeDecimalPlacesOfTheTradingSizeExceededTheLimit order endpoint::The decimal places of the trading size exceeded the limit
	ErrCodeDecimalPlacesOfTheTradingSizeExceededTheLimit = 33029

	// ErrCodeWithdrawalSuspended withdrawal endpoint: account suspended
	ErrCodeWithdrawalSuspended = 34001
	// ErrCodeNoWithdrawalAddress withdrawal endpoint: address required
	ErrCodeNoWithdrawalAddress = 34002
	// ErrCodeTokenCannotBeWithdrawnToXxAtTheMoment withdrawal endpoint: incorrect address
	ErrCodeTokenCannotBeWithdrawnToXxAtTheMoment = 34003
	// ErrCodeWithdrawalFeeIsSmallerThanMinimumLimit withdrawal endpoint: incorrect fee
	ErrCodeWithdrawalFeeIsSmallerThanMinimumLimit = 34004
	// ErrCodeWithdrawalFeeExceedsTheMaximumLimit withdrawal endpoint: incorrect withdrawal fee
	ErrCodeWithdrawalFeeExceedsTheMaximumLimit = 34005
	// ErrCodeWithdrawalAmountIsLowerThanTheMinimumLimit minimum withdrawal amount%} endpoint: incorrect amount
	ErrCodeWithdrawalAmountIsLowerThanTheMinimumLimit = 34006
	// ErrCodeWithdrawalAmountExceedsTheMaximumLimit maximum withdrawal amount endpoint: incorrect amount
	ErrCodeWithdrawalAmountExceedsTheMaximumLimit = 34007
	// ErrCodeWithdrawalInsufficientBalance transfer & withdrawal endpoint: insufficient balance
	ErrCodeWithdrawalInsufficientBalance = 34008
	// ErrCodeWithdrawalAmountExceedsTheDailyLimit withdrawal endpoint: withdrawal limit exceeded
	ErrCodeWithdrawalAmountExceedsTheDailyLimit = 34009
	// ErrCodeTransferAmountMustBeLargerThanZero transfer endpoint: incorrect amount
	ErrCodeTransferAmountMustBeLargerThanZero = 34010
	// ErrCodeConditionsNotMet transfer & withdrawal endpoint: conditions not met, e.g. KYC level
	ErrCodeConditionsNotMet = 34011
	// ErrCodeMinimumWithdrawalAmountForNeoIsOneAndTheAmountMustBeAnInteger withdrawal endpoint: special requirements
	ErrCodeMinimumWithdrawalAmountForNeoIsOneAndTheAmountMustBeAnInteger = 34012
	// ErrCodeTransferNoInstrumentID transfer endpoint: Token margin trading instrument ID required
	ErrCodeTransferNoInstrumentID = 34013
	// ErrCodeTransferLimited transfer endpoint：Transfer limited
	ErrCodeTransferLimited = 34014
	// ErrCodeSubaccountNotExist transfer endpoint: subaccount does not exist
	ErrCodeSubaccountNotExist = 34015
	// ErrCodeTransferSuspended transfer endpoint: either end of the account does not authorize the transfer
	ErrCodeTransferSuspended = 34016
	// ErrCodeTransferAccountSuspended transfer & withdrawal endpoint: either end of the account does not authorize the transfer
	ErrCodeTransferAccountSuspended = 34017
	// ErrCodeIncorrectTradesPassword incorrect trades password
	ErrCodeIncorrectTradesPassword = 34018
	// ErrCodeNotBindEmailBeforeWithdrawal withdrawal endpoint : email required
	ErrCodeNotBindEmailBeforeWithdrawal = 34019
	// ErrCodeNotBindFundsPasswordBeforeWithdrawal withdrawal endpoint : funds password required
	ErrCodeNotBindFundsPasswordBeforeWithdrawal = 34020
	// ErrCodeNotVerifiedAddress withdrawal endpoint
	ErrCodeNotVerifiedAddress = 34021
	// ErrCodeWithdrawalsAreNotAvailableForSubAccounts withdrawal endpoint
	ErrCodeWithdrawalsAreNotAvailableForSubAccounts = 34022
)
