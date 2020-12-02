# hexbank
Banking Service following hexagonal architecture


Customer
        Create Account
        Deposit (Cash, CQ, Transfer)
        Withdrawn
        Balance
                                                            Bank
                            
 API (http REST)
    MUX
    handlerfunc(){
        AccountCreator
            CreateAccount()
        
        AccountManager
            ?
            
        FundsManager        
            DepositFunds()
            WithdrawFunds()
            RetrieveBalance()
 
 GRPC
        AccountCreator
             CreateAccount()
         
         AccountManager
             ?
             
         FundsManager        
             DepositFunds()
             WithdrawFunds()
             RetrieveBalance()
 
 Tester
        Account                    