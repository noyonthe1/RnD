package main


import (
	"time"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ProxyChaincode example simple Chaincode implementation
type ProxyChaincode struct {
}

type Institute struct {
	Id      int `json:"id"`
	Name       string `json:"name"`  
	Address       string    `json:"address"`
	ContactNumber      string `json:"contactNumber"`
}

type Campaign struct {
	Id      int `json:"id"`
	InstitutionId      int `json:"institutionId"`
	Name       string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate  time.Time `json:"endDate"`  
	Status       string    `json:"status"`
	Properties      string `json:"properties"`
	Institute Institute
}

type Proposal struct {
	Id      int `json:"id"`
	CampaignId      int `json:"campaignId"`
	ProposalDetail       string    `json:"proposalDetail"`
	Option       string    `json:"option"`
	isDirectorSelection      string `json:"isDirectorSelection"`
	Campaign Campaign
}

type ShareHolder struct {
	Id      int `json:"id"`
	CampaignId      int `json:"campaignId"`
	ControlNumber       string    `json:"controlNumber"`
	NoOfVoteRepresents       string    `json:"voteRepresent"`
	Campaign Campaign
}

type Vote struct {
	Id      int `json:"id"`
	ProposalId      int `json:"proposalId"`
	Answer       string    `json:"answer"`
	ControlNumber       string    `json:"controlNumber"`
	Proposal Proposal
}

type VoteCount struct {
	Id      int `json:"id"`
	CampaignId int`json:"campaignId"`
	ProposalId int`json:"proposalId"`
	VoteId      int `json:"voteId"`
	ForVoteCount       int    `json:"forVoteCount"`
	AgainstVoteCount       int    `json:"againstVoteCount"`
	AbstainVoteCount       int    `json:"abstainVoteCount"`
}


func (t *ProxyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	return shim.Success(nil)
}

func (t *ProxyChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
		return shim.Error("Unknown supported call")
}

// Transaction makes payment of X units from A to B
func (t *ProxyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
        
	function, args := stub.GetFunctionAndParameters()
	
    if function == "CreateInstitute" { 
		return t.CreateInstitute(stub, args)
	} else if function == "GetInstituteById" { 
		return t.GetInstituteById(stub, args)
	} else if function == "UpdateInstitute" { 
		return t.UpdateInstitute(stub, args)
	} else if function == "CreateCampaign" { 
		return t.CreateCampaign(stub, args)
	} else if function == "UpdateCampaign" { 
		return t.UpdateCampaign(stub, args)
	} else if function == "GetCampaignById" { 
		return t.GetCampaignById(stub, args)
	} else if function == "CreateProposal" { 
		return t.CreateProposal(stub, args)
	} else if function == "UpdateProposal" { 
		return t.UpdateProposal(stub, args)
	} else if function == "GetProposalById" { 
		return t.GetProposalById(stub, args)
	} else if function == "CreateShareHolder" { 
		return t.CreateShareHolder(stub, args)
	} else if function == "UpdateShareHolder" { 
		return t.UpdateShareHolder(stub, args)
	} else if function == "GetShareHolderById" { 
		return t.GetShareHolderById(stub, args)
	} else if function == "CreateVote" { 
		return t.CreateVote(stub, args)
	} else if function == "UpdateVote" { 
		return t.UpdateVote(stub, args)
	} else if function == "GetVoteById" { 
		return t.GetVoteById(stub, args)
	} else if function == "CreateVoteCount" { 
		return t.CreateVoteCount(stub, args)
	} else if function == "GetVoteCountById" { 
		return t.GetVoteCountById(stub, args)
	} else if function == "getRecordsByRange" { 
		return t.getRecordsByRange(stub, args)
	}
    fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *ProxyChaincode) CreateInstitute(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fmt.Println("- start init institution")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}

	
	Id,_ := strconv.Atoi(args[0])
	Name := args[1]
	Address := args[2]
	ContactNumber := args[3]


    
	// ==== Create institution object and marshal to JSON ====

	objectType := "institution_" + (args[0]);
	objInstitute := &Institute{Id, Name, Address, ContactNumber}
	

    //compositeKey1, _ := stub.CreateCompositeKey("institution", []string{objInstitute.Name, objInstitute.Address})
	//objInstitute.Name = compositeKey1 

	instituteJSONasBytes, err := json.Marshal(objInstitute)
	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save institute to state ===

	err = stub.PutState(objectType, instituteJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) UpdateInstitute(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fmt.Println("- start init institution")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}

	objectType := "institution_" + args[0];
    marbleAsBytes,err1 := stub.GetState(objectType);
	if err1 != nil {
		return shim.Error(err.Error())
	}else if marbleAsBytes == nil {
		fmt.Println("institute does not exists ")
		return shim.Error("institute does not exists.")
	}
	


	Id,_ := strconv.Atoi(args[0])
	Name := args[1]
	Address := args[2]
	ContactNumber := args[3]


	// ==== Create institution object and marshal to JSON ====


	objInstitute := &Institute{Id, Name, Address, ContactNumber}
	instituteJSONasBytes, err := json.Marshal(objInstitute)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save institute to state ===

	err = stub.PutState(objectType, instituteJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) GetInstituteById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "institution_" + args[0];
    instituteAsBytes,err := stub.GetState(objectType);
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(instituteAsBytes)

}

func (t *ProxyChaincode) GetInstituteObjById(stub shim.ChaincodeStubInterface,instituteId string) Institute {
	objectType := "institution_" +instituteId;
    instituteAsBytes,err := stub.GetState(objectType);
	newInstitute := Institute{} 
	if err != nil {
		return newInstitute
	}
	err = json.Unmarshal(instituteAsBytes, &newInstitute)

	return newInstitute
}
//============================= Campaign =======================================
func (t *ProxyChaincode) CreateCampaign(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	fmt.Println("- start creating campaign")
   
    if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	timeLayOut := "2006-01-02"
	Id,_ := strconv.Atoi(args[0])
	InstitutionId,_ :=  strconv.Atoi(args[1])
	Name := args[2]
	StartDate,_ := time.Parse(timeLayOut, args[3])
	EndDate,_ := time.Parse(timeLayOut, args[4])
	Status := args[5]
	Properties := args[6]

	// ==== Create institution object and marshal to JSON ====

	objectType := "campaign_" + (args[0]);
	institute := Institute{} 
	objCampaign := &Campaign{Id,InstitutionId, Name, StartDate,EndDate,Status, Properties,institute}
	campaignJSONasBytes, err := json.Marshal(objCampaign)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save Campaign to state ===

	err = stub.PutState(objectType, campaignJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) UpdateCampaign(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	fmt.Println("- start init campaign")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	objectType := "campaign_" + args[0];
    campaignAsBytes,err1 := stub.GetState(objectType);
	if err1 != nil {
		return shim.Error(err.Error())
	}else if campaignAsBytes == nil {
		fmt.Println("campaign does not exists ")
		return shim.Error("campaign does not exists.")
	}
	


	timeLayOut := "2006-01-02"
	Id,_ := strconv.Atoi(args[0])
	InstitutionId,_ :=  strconv.Atoi(args[1])
	Name := args[2]
	StartDate,_ := time.Parse(timeLayOut, args[3])
	EndDate,_ := time.Parse(timeLayOut, args[4])
	Status := args[5]
	Properties := args[6]


	// ==== Create campaign object and marshal to JSON ====

    institute := Institute{} 
	objCampaign :=&Campaign{Id,InstitutionId, Name, StartDate,EndDate,Status, Properties,institute}
	campaignJSONasBytes, err := json.Marshal(objCampaign)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save campaign to state ===

	err = stub.PutState(objectType, campaignJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) GetCampaignById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "campaign_" + args[0];
    campaignAsBytes,err := stub.GetState(objectType);

	//Institute := Institute{}
	
    campaign := Campaign{} 
	err = json.Unmarshal(campaignAsBytes, &campaign)
	if err != nil {

		return shim.Error(err.Error())

	}
    insId := strconv.Itoa(campaign.InstitutionId)
    campaign.Institute  = t.GetInstituteObjById(stub,insId)
    campaignJSONasBytes,err := json.Marshal(campaign)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(campaignJSONasBytes)

}

func (t *ProxyChaincode) GetCampaignObjById(stub shim.ChaincodeStubInterface, campaignId string) Campaign {

    
	objectType := "campaign_" + campaignId;
    campaignAsBytes,err := stub.GetState(objectType);
	
    campaign := Campaign{} 
	err = json.Unmarshal(campaignAsBytes, &campaign)
	if err != nil {

		return campaign

	}
    insId := strconv.Itoa(campaign.InstitutionId)
    campaign.Institute  = t.GetInstituteObjById(stub,insId)
   
	return campaign
}
//============================= Proposal =======================================

func (t *ProxyChaincode) CreateProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start creating Proposal")
   
    if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	
	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ :=  strconv.Atoi(args[1])
	ProposalDetail := args[2]
	Option := args[3]
	isDirectorSelection := args[4]
    campaign := Campaign{} 
	// ==== Create institution object and marshal to JSON ====

	objectType := "proposal_" + (args[0]);
	objProposal := &Proposal{Id,CampaignId, ProposalDetail, Option,isDirectorSelection,campaign}
	ProposalJSONasBytes, err := json.Marshal(objProposal)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save Proposal to state ===

	err = stub.PutState(objectType, ProposalJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) UpdateProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start init Proposal")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	objectType := "proposal_" + args[0];
    ProposalAsBytes,err1 := stub.GetState(objectType);
	if err1 != nil {
		return shim.Error(err.Error())
	}else if ProposalAsBytes == nil {
		fmt.Println("Proposal does not exists ")
		return shim.Error("Proposal does not exists.")
	}
	


	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ :=  strconv.Atoi(args[1])
	ProposalDetail := args[2]
	Option := args[3]
	isDirectorSelection := args[4]


	// ==== Create Proposal object and marshal to JSON ====

    campaign := Campaign{} 
	objProposal :=&Proposal{Id,CampaignId, ProposalDetail, Option,isDirectorSelection,campaign}
	proposalJSONasBytes, err := json.Marshal(objProposal)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save Proposal to state ===

	err = stub.PutState(objectType, proposalJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) GetProposalById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "proposal_" + args[0];
    proposalAsBytes,err := stub.GetState(objectType);
	if err != nil {
		return shim.Error(err.Error())
	}
	proposal := Proposal{}
	err = json.Unmarshal(proposalAsBytes, &proposal)
	if err != nil {
		return shim.Error(err.Error())
	}
	cmpId := strconv.Itoa(proposal.CampaignId)
	proposal.Campaign =  t.GetCampaignObjById(stub,cmpId)

	proposalJSONasBytes,err := json.Marshal(proposal)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(proposalJSONasBytes)

}
func (t *ProxyChaincode) GetProposalObjById(stub shim.ChaincodeStubInterface, proposalId string) Proposal {

   
	objectType := "proposal_" + proposalId;
    proposalAsBytes,err := stub.GetState(objectType);
	
	proposal := Proposal{}
	err = json.Unmarshal(proposalAsBytes, &proposal)
	if err != nil {
		return proposal
	}
	cmpId := strconv.Itoa(proposal.CampaignId)
	proposal.Campaign =  t.GetCampaignObjById(stub,cmpId)

	return proposal
}

//============================= ShareHolder =====================================


func (t *ProxyChaincode) CreateShareHolder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fmt.Println("- start creating ShareHolder")
   
    if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	

	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ :=  strconv.Atoi(args[1])
	ControlNumber := args[2]
	NoOfVoteRepresents := args[3]
    campaign := Campaign{} 
	// ==== Create institution object and marshal to JSON ====

	objectType := "shareHolder_" + (args[0]);
	objShareHolder := &ShareHolder{Id,CampaignId, ControlNumber, NoOfVoteRepresents,campaign}
	shareHolderJSONasBytes, err := json.Marshal(objShareHolder)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save ShareHolder to state ===

	err = stub.PutState(objectType, shareHolderJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) UpdateShareHolder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start init ShareHolder")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	objectType := "shareHolder_" + args[0];
    shareHolderAsBytes,err1 := stub.GetState(objectType);
	if err1 != nil {
		return shim.Error(err.Error())
	}else if shareHolderAsBytes == nil {
		fmt.Println("ShareHolder does not exists ")
		return shim.Error("ShareHolder does not exists.")
	}
	
	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ :=  strconv.Atoi(args[1])
	ControlNumber := args[2]
	NoOfVoteRepresents := args[3]


	// ==== Create ShareHolder object and marshal to JSON ====
    campaign := Campaign{}

	objShareHolder :=&ShareHolder{Id,CampaignId, ControlNumber, NoOfVoteRepresents,campaign}
	shareHolderJSONasBytes, err := json.Marshal(objShareHolder)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save ShareHolder to state ===

	err = stub.PutState(objectType, shareHolderJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) GetShareHolderById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "shareHolder_" + args[0];
    shareHolderAsBytes,err := stub.GetState(objectType);
	if err != nil {
		return shim.Error(err.Error())
	}

	shareHolder := ShareHolder{}
	err = json.Unmarshal(shareHolderAsBytes, &shareHolder)
	if err != nil {
		return shim.Error(err.Error())
	}
	cmpId := strconv.Itoa(shareHolder.CampaignId)
	shareHolder.Campaign =  t.GetCampaignObjById(stub,cmpId)

    shareHolderJSONasBytes,err := json.Marshal(shareHolder)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(shareHolderJSONasBytes)
}
//============================= Vote =====================================


func (t *ProxyChaincode) CreateVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fmt.Println("- start creating Vote")
   
    if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	

	Id,_ := strconv.Atoi(args[0])
	ProposalId,_ :=  strconv.Atoi(args[1])
	Answer := args[2]
	ControlNumber := args[3]


	
    
	proposal := Proposal{}
	objVote := &Vote{Id,ProposalId, Answer, ControlNumber,proposal}
	voteJSONasBytes, err := json.Marshal(objVote)

	if err != nil {
		return shim.Error(err.Error())
	}


	// ==== Create institution object and marshal to JSON ====
	
   newProposal :=  t.GetProposalObjById(stub,args[1])
   CampaignId :=  strconv.Itoa(newProposal.Campaign.Id)

   objectType := "voteCount_" + CampaignId + "_" + args[1] ;

   newVoteCount := VoteCount{} 

   voteCountJson,err2 := stub.GetState(objectType);
	if err2 != nil {
		return shim.Error(err.Error())
	}else if voteCountJson == nil {
		fmt.Println("no votecount for this proposal 1")
		
        newVoteCount.Id = 1
		newVoteCount.CampaignId = newProposal.Campaign.Id
		newVoteCount.ProposalId = newProposal.Id
		newVoteCount.VoteId = 1
		newVoteCount.ForVoteCount = 0;
		newVoteCount.AgainstVoteCount = 0 ;
		newVoteCount.AbstainVoteCount = 0;
	}else {
   		err3 := json.Unmarshal(voteCountJson, &newVoteCount)
	   	if err3 != nil {
			return shim.Error(err3.Error())
		}
	}
	
	//
	objectType = "vote_" + args[1] + "_" + args[3];
    newVote := Vote{} 
    voteJson,err3 := stub.GetState(objectType);
	if err3 != nil {
		return shim.Error(err.Error())
	}else if voteJson != nil {
		err4 := json.Unmarshal(voteJson, &newVote)
		if err4 != nil {
			return shim.Error(err4.Error())
		}
		
		if newVote.Answer == "1"{
		newVoteCount.ForVoteCount -= 1
		}else if newVote.Answer == "2"{
		newVoteCount.AgainstVoteCount -= 1
		}else if newVote.Answer == "3"{
		newVoteCount.AbstainVoteCount -= 1
		}	
	}
	//

	// === Save Vote to state ===

	err = stub.PutState(objectType, voteJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}


	votecount_id := strconv.Itoa(newVoteCount.Id)
	votecount_CampaignId := strconv.Itoa(newVoteCount.CampaignId)
	votecount_ProposalId := strconv.Itoa(newVoteCount.ProposalId)
	votecount_VoteId :=  strconv.Itoa(newVoteCount.VoteId)

	if Answer == "1"{
		newVoteCount.ForVoteCount += 1
	}else if Answer == "2"{
		newVoteCount.AgainstVoteCount += 1
	}else if Answer == "3"{
		newVoteCount.AbstainVoteCount += 1
	}

	votecount_ForVoteCount :=  strconv.Itoa(newVoteCount.ForVoteCount)
	votecount_AgainstVoteCount :=  strconv.Itoa(newVoteCount.AgainstVoteCount)
	votecount_AbstainVoteCount :=  strconv.Itoa(newVoteCount.AbstainVoteCount)

	args2 := []string{votecount_id,votecount_CampaignId, votecount_ProposalId, votecount_VoteId,votecount_ForVoteCount,votecount_AgainstVoteCount,votecount_AbstainVoteCount}
	result :=  t.CreateVoteObjCount(stub,args2)
	
	fmt.Println(result);
	
	fmt.Println("==============================================");
	return shim.Success(nil)
}

func (t *ProxyChaincode) UpdateVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start init Vote")
	if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	objectType := "vote_" + args[0];
    voteAsBytes,err1 := stub.GetState(objectType);
	if err1 != nil {
		return shim.Error(err.Error())
	}else if voteAsBytes == nil {
		fmt.Println("Vote does not exists ")
		return shim.Error("Vote does not exists.")
	}
	
	Id,_ := strconv.Atoi(args[0])
	ProposalId,_ :=  strconv.Atoi(args[1])
	Answer := args[2]
	ControlNumber := args[3]


	// ==== Create Vote object and marshal to JSON ====

    proposal := Proposal{}
	objVote :=&Vote{Id,ProposalId, Answer, ControlNumber,proposal}
	VoteJSONasBytes, err := json.Marshal(objVote)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save Vote to state ===

	err = stub.PutState(objectType, VoteJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) GetVoteById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	objectType := "vote_" + args[0] + "_" +args[1];
    VoteAsBytes,err := stub.GetState(objectType);
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(VoteAsBytes)
}
//============================= VoteCount =====================================


func (t *ProxyChaincode) CreateVoteCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	fmt.Println("- start creating VoteCount")
   
    if len(args[0]) <= 0 { 
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
    if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	
	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ := strconv.Atoi(args[1])
	ProposalId,_ := strconv.Atoi(args[2])
	VoteId,_ :=  strconv.Atoi(args[3])
	ForVoteCount,_ :=  strconv.Atoi(args[4])
	AgainstVoteCount,_ :=  strconv.Atoi(args[5])
	AbstainVoteCount,_ :=  strconv.Atoi(args[6])

	// ==== Create VoteCount object and marshal to JSON ====

	

	// ==== Create votecount object and marshal to JSON ====

	objectType := "voteCount_" + args[1] + "_" + args[2];
	objVoteCount := &VoteCount{Id, CampaignId, ProposalId, VoteId, ForVoteCount, AgainstVoteCount , AbstainVoteCount}
	
	voteCountJSONasBytes, err := json.Marshal(objVoteCount)

	if err != nil {
		return shim.Error(err.Error())
	}
	// === Save VoteCount to state ===

	err = stub.PutState(objectType, voteCountJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (t *ProxyChaincode) CreateVoteObjCount(stub shim.ChaincodeStubInterface, args []string) *VoteCount {

	var err error
	
	Id,_ := strconv.Atoi(args[0])
	CampaignId,_ := strconv.Atoi(args[1])
	ProposalId,_ := strconv.Atoi(args[2])
	VoteId,_ :=  strconv.Atoi(args[3])
	ForVoteCount,_ :=  strconv.Atoi(args[4])
	AgainstVoteCount,_ :=  strconv.Atoi(args[5])
	AbstainVoteCount,_ :=  strconv.Atoi(args[6])

	// ==== Create votecount object and marshal to JSON ====

	objectType := "voteCount_" + args[1] + "_" + args[2];
	objVoteCount := &VoteCount{Id, CampaignId, ProposalId, VoteId, ForVoteCount, AgainstVoteCount , AbstainVoteCount}
	
	voteCountJSONasBytes, err := json.Marshal(objVoteCount)

	err = stub.PutState(objectType, voteCountJSONasBytes)

	if err != nil {
		return objVoteCount
	}
	return objVoteCount

}

func (t *ProxyChaincode) GetVoteCountById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	objectType := "voteCount_" + args[0];
    VoteCountAsBytes,err := stub.GetState(objectType);
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(VoteCountAsBytes)
}

func (t *ProxyChaincode) getRecordsByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)

	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	var buffer bytes.Buffer

	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")

		// Record is a JSON object, so we write as-is

		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}


func main() {
	err := shim.Start(new(ProxyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
