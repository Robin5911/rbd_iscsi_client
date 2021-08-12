package rbd_iscsi_client

import (
	"encoding/json"
	"fmt"
	"github.com/TonyZhang1989/rbd_iscsi_client/models"
	"log"
	"net/url"
)
type BasicAuth struct {
	Username string
	Password string
}
type RBDISCSIClient struct {
	ApiUrl string
	Timeout int
	Secure bool
	Auth BasicAuth
}
func (rbdIscsiClient *RBDISCSIClient) GetApi() string{
	apiUrl := fmt.Sprintf("%s/api",rbdIscsiClient.ApiUrl)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil || (httpCode%200) > 2   {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
	return ""
}
func(rbdIscsiClient *RBDISCSIClient) GetConfig() (models.Config,int,error){
	apiUrl := fmt.Sprintf("%s/api/config",rbdIscsiClient.ApiUrl)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil {
		return models.Config{},httpCode, err
	}
	fmt.Println(string(resp))
	var config models.Config
	err = json.Unmarshal(resp,&config)
	if err != nil {
		log.Println("json unmarshal err : ",err.Error())
		return models.Config{}, httpCode,err
	}
	return config,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) GetSysInfo(t string) string{
	/*
	Get system info of <type>
	 */
	apiUrl := fmt.Sprintf("%s/api/sysinfo/%s",rbdIscsiClient.ApiUrl,t)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil || (httpCode%200) > 2 {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
	return ""
}
func(rbdIscsiClient *RBDISCSIClient) GetGatewayinfo() interface{}{
	apiUrl := fmt.Sprintf("%s/api/gatewayinfo",rbdIscsiClient.ApiUrl)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil || (httpCode%200) > 2 {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
	return ""
}
func(rbdIscsiClient *RBDISCSIClient) GetTargets() interface{}{
	apiUrl := fmt.Sprintf("%s/api/targets",rbdIscsiClient.ApiUrl)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil || (httpCode%200) > 2 {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
	return ""
}
func(rbdIscsiClient *RBDISCSIClient) GetTargetInfo(targetIqn string,mode,controls interface{}) (models.GenericResp,int,error){
	apiUrl := fmt.Sprintf("%s/api/targetinfo/%s",rbdIscsiClient.ApiUrl, targetIqn)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return models.GenericResp{},httpCode,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func (rbdIscsiClient *RBDISCSIClient) CreateTargetIqn(targetIqn string,mode,controls interface{}) (models.GenericResp,int,error){
	/*
		Create a target iqn from the gateways
	*/
	apiUrl := fmt.Sprintf("%s/api/target/%s",rbdIscsiClient.ApiUrl, targetIqn)
	payload := url.Values{}
	if mode != nil {
		payload.Add("mode",mode.(string))
	}
	if controls != nil {
		payload.Add("controls",controls.(string))
	}
	resp,httpCode,err := Request(apiUrl,"PUT",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return models.GenericResp{},httpCode,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) DeleteTargetIqn(targetIqn string){
	/*
	Delete a target iqn from the gateways
	 */
	apiUrl := fmt.Sprintf("%s/api/target/%s",rbdIscsiClient.ApiUrl, targetIqn)
	resp,httpCode,err := Request(apiUrl,"DELETE",nil,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return
	}
	fmt.Println(string(resp))
}
func(rbdIscsiClient *RBDISCSIClient) GetClients(targetIqn string){
	/*
	List clients defined to the configuration
	 */
	apiUrl := fmt.Sprintf("%s/api/client/%s", rbdIscsiClient.ApiUrl,targetIqn)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return
	}
	fmt.Println(string(resp))
}
func(rbdIscsiClient *RBDISCSIClient) GetClientInfo(targetIqn, clientIqn string){
	/*
	Fetch the Client information from the gateway
	 */
	apiUrl := fmt.Sprintf("%s/api/client/%s/%s", rbdIscsiClient.ApiUrl,targetIqn, clientIqn)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return
	}
	fmt.Println(string(resp))
}
func(rbdIscsiClient *RBDISCSIClient) CreateClient(targetIqn, clientIqn string) (models.GenericResp,int,error){
	/*
	Create a client
	 */
	apiUrl := fmt.Sprintf("%s/api/client/%s/%s", rbdIscsiClient.ApiUrl,targetIqn, clientIqn)
	resp,httpCode,err := Request(apiUrl,"PUT",nil,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) DeleteClient(targetIqn, clientIqn string) (models.GenericResp,int,error){
	/*
	Delete a client
	 */
	apiUrl := fmt.Sprintf("%s/api/client/%s/%s", rbdIscsiClient.ApiUrl,targetIqn, clientIqn)
	resp,httpCode,err := Request(apiUrl,"DELETE",nil,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) SetClientAuth(targetIqn, clientIqn string,auth models.ClientAuth) (models.GenericResp,int,error){
	/*
	Set the client chap credentials
	 */
	apiUrl := fmt.Sprintf("%s/api/clientauth/%s/%s", rbdIscsiClient.ApiUrl,targetIqn, clientIqn)
	payload := url.Values{}
	payload.Add("username",auth.Username)
	payload.Add("password",auth.Password)
	resp,httpCode,err := Request(apiUrl,"PUT",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) GetDisks(){
	apiUrl := fmt.Sprintf("%s/api/disks",rbdIscsiClient.ApiUrl)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return
	}
	if (httpCode%200) > 2 {
		log.Println("api resp http code : ",httpCode)
		return
	}
	fmt.Println(string(resp))
}
func(rbdIscsiClient *RBDISCSIClient) CreateDisk(pool,image string,size,extras interface{}) (models.GenericResp,int,error){
	/*
	Add a disk to the gateway
	 */
	apiUrl := fmt.Sprintf("%s/api/disk/%s/%s",rbdIscsiClient.ApiUrl,pool,image)
	payload := url.Values{}
	payload.Add("mode","create")
	payload.Add("pool",pool)
	payload.Add("image",image)
	resp,httpCode,err := Request(apiUrl,"PUT",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil


}
func(rbdIscsiClient *RBDISCSIClient) findDisk(pool,image string){
	apiUrl := fmt.Sprintf("/api/disk/%s/%s",pool,image)
	resp,httpCode,err := Request(apiUrl,"GET",nil,nil,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil || (httpCode%200) > 2 {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
}
func(rbdIscsiClient *RBDISCSIClient) DeleteDisk(pool,image string) (models.GenericResp,int,error){
	/*
	Delete a disk definition from the gateway
	By default it will not delete the rbd image from the pool.
    If preserve_image is set to True, then this only tells the
    gateway to forget about this volume.   This is typically done
    when the volume isn't used as an export anymore.
	 */
	apiUrl := fmt.Sprintf("%s/api/disk/%s/%s",rbdIscsiClient.ApiUrl,pool,image)
	resp,httpCode,err := Request(apiUrl,"DELETE",nil,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) RegisterDisk(targetIqn,volume string) (models.GenericResp,int,error){
	/*
	Add the voluem to the target definition
	This is done after the disk is created in a pool, and
	before the disk can be exported to an initiator.
	 */
	apiUrl := fmt.Sprintf("%s/api/targetlun/%s",rbdIscsiClient.ApiUrl,targetIqn)
	payload := url.Values{}
	payload.Add("disk",volume)
	resp,httpCode,err := Request(apiUrl,"PUT",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) UnregisterDisk(targetIqn,volume string) (models.GenericResp,int,error){
	/*
	Remove the volume from the target definition.
	This is done after the disk is unexported from an initiator
	and before the disk can be deleted from the gateway.
	 */
	apiUrl := fmt.Sprintf("%s/api/targetlun/%s",rbdIscsiClient.ApiUrl,targetIqn)
	payload := url.Values{}
	payload.Add("disk",volume)
	resp,httpCode,err := Request(apiUrl,"DELETE",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
func(rbdIscsiClient *RBDISCSIClient) ExportDisk(targetIqn,clientIqn,pool,disk string)  (models.GenericResp,int,error){
	/*
	Add a disk to export to a client.
	 */
	apiUrl := fmt.Sprintf("%s/api/clientlun/%s/%s",rbdIscsiClient.ApiUrl,targetIqn,clientIqn)
	payload := url.Values{}
	payload.Add("disk",fmt.Sprintf("%s/%s",pool,disk))
	payload.Add("client_iqn",clientIqn)
	resp,httpCode,err := Request(apiUrl,"PUT",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil

}
func(rbdIscsiClient *RBDISCSIClient) UnexportDisk(targetIqn,clientIqn,pool,disk string) (models.GenericResp,int,error){
	/*
	Remove a disk to export to a client
	 */
	apiUrl := fmt.Sprintf("%s/api/clientlun/%s/%s",rbdIscsiClient.ApiUrl,targetIqn,clientIqn)
	payload := url.Values{}
	payload.Add("disk",fmt.Sprintf("%s/%s",pool,disk))
	resp,httpCode,err := Request(apiUrl,"DELETE",payload,FormContentType,rbdIscsiClient.Auth,rbdIscsiClient.Timeout)
	if err != nil  {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},-1,err
	}
	var genericResp models.GenericResp
	err = json.Unmarshal(resp,&genericResp)
	if err != nil {
		log.Println("api resp err : ",err.Error())
		return models.GenericResp{},httpCode,err
	}
	return genericResp,httpCode,nil
}
