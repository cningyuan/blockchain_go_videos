package BLC

import (
	"bytes"
	"log"
	"encoding/gob"
	"encoding/hex"
)

func handleVersion(request []byte,bc *Blockchain)  {

	var buff bytes.Buffer
	var payload Version

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	//Version
	//1. Version
	//2. BestHeight
	//3. 节点地址

	bestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if bestHeight > foreignerBestHeight {
		sendVersion(payload.AddrFrom,bc)
	} else if bestHeight < foreignerBestHeight {
		// 去向主节点要信息
		sendGetBlocks(payload.AddrFrom)
	}


}

func handleAddr(request []byte,bc *Blockchain)  {




}

func handleGetblocks(request []byte,bc *Blockchain)  {


	var buff bytes.Buffer
	var payload GetBlocks

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := bc.GetBlockHashes()

	sendInv(payload.AddrFrom, BLOCK_TYPE, blocks)


}

func handleGetData(request []byte,bc *Blockchain)  {

	var buff bytes.Buffer
	var payload GetData

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == BLOCK_TYPE {

		block, err := bc.GetBlock([]byte(payload.Hash))
		if err != nil {
			return
		}

		sendBlock(payload.AddrFrom, &block)
	}

	if payload.Type == "tx" {

	}


}

func handleBlock(request []byte,bc *Blockchain)  {




}

func handleTx(request []byte,bc *Blockchain)  {

}


func handleInv(request []byte,bc *Blockchain)  {

	var buff bytes.Buffer
	var payload Inv

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	// Ivn 3000 block hashes [][]

	if payload.Type == BLOCK_TYPE {

		blockHash := payload.Items[0]
		sendGetData(payload.AddrFrom, BLOCK_TYPE , blockHash)

	}

	if payload.Type == TX_TYPE {

	}

}