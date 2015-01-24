package gostratum

import (
    "net";
    "bufio";
    "time";
    "encoding/json";
    "strconv"
)


type Client struct
{
    socket net.Conn
    seq uint64
    dispatcher *Dispatcher
    decoder *Decoder
    encoder *Encoder
    timeout time.Duration
}

func Connect(host string) (*Client, error){

    var client Client
    var err error
    client.socket, err = net.Dial("tcp", host)
    if err!=nil {return nil,err}
    client.seq = 0
    client.dispatcher = MakeDispatcher()
    client.decoder = MakeDecoder()
    client.encoder = MakeEncoder()
    client.SetTimeout(10)
    go client.Listen()
    return &client, nil
}

func (c *Client) SetTimeout(timeout int64){
    c.timeout = time.Duration(timeout) * time.Second
}

func (c *Client) Listen(){
    c.socket.SetReadDeadline(time.Time{})
    for{
        result, err := bufio.NewReader(c.socket).ReadString('\n')
        if err!=nil{
            c.dispatcher.Error(err)
            continue
        }
        response := &Response{}
        err = c.decoder.Decode(result, response)
        if err!=nil{
            c.dispatcher.Error(err)
            continue
        }
        c.dispatcher.Process(response)
    }
}

func (c *Client) Send(request *Request) *Response{
    c.seq++
    request.ID = c.seq
    msg, err:= c.encoder.Encode(request)
    if err!=nil {return &Response{Error:err}}
    action,err := c.dispatcher.RegisterRequest(request)
    if err != nil {return &Response{Error:err}}
    _, err = c.socket.Write([]byte(msg));
    if err!=nil{
        c.dispatcher.Cancel(request.ID)
        return &Response{Error:err}
    }
    action.SetTimeout(c.timeout, func(){
        c.dispatcher.Cancel(request.ID)
    })
    return action.Wait()
}

func (c *Client) Request(method string, params ...string) *Response{
    return c.Send(&Request{Method:method, Params:params});
}


func (c *Client) ServerVersion() (string, error){
    return c.decoder.DecodeStringResult(c.Request("server.version"))
}

func (c *Client) ServerBanner() (string, error){
    return c.decoder.DecodeStringResult(c.Request("server.version"))
}

func (c *Client) ServerDontationAddress() (string, error){
    return c.decoder.DecodeStringResult(c.Request("server.version"))
}

func (c *Client) Subscribe(method string, eventHandler func(*json.RawMessage), msgHandler func(*json.RawMessage), params ... string) error{
    c.dispatcher.RegisterNotifiactionHandler(method, eventHandler)
    response := c.Send(&Request{Method:method, Params:params})
    if response.Error !=nil {return response.Error}
    if msgHandler != nil{
        msgHandler(response.Result)
    }else{
        eventHandler(response.Result)
    }
    return nil
}


func (c *Client) PeersSubscribe(callback func([]Peer, error)) error{
    return c.Subscribe("server.peers.subscribe", WrapPeersHandler(callback, c.decoder), nil)
}

func (c *Client) BlockHeaderSubscribe(callback func([]BlockHeader, error)) error{
    return c.Subscribe("blockchain.headers.subscribe", WrapBlockHeadersHandler(callback, c.decoder), nil)
}

func (c *Client) NumBlocksSubscribe(callback func(int, error))error{
    return c.Subscribe("blockchain.numblocks.subscribe", WrapNumBlocksHandler(callback, c.decoder), nil)
}

func (c *Client) AddressSubscribe(address string, callback func(string, string, error)) error {
    return c.Subscribe("blockchain.address.subscribe", WrapAddressHandler(callback, c.decoder), func(result *json.RawMessage){
        status, err := c.decoder.DecodeString(result)
        callback(address, status, err)
    }, address)
}


func (c *Client) AddressGetHistory(address string) ([]AddressTransaction, error){
    return c.decoder.DecodeAddressTransactionsResult(c.Request("blockchain.address.get_history", address))
}

func (c *Client) AddressGetMemPool(address string) error{
    response := c.Request("blockchain.address.get_mempool", address)
    return nil
}

func (c *Client) AddressGetBalance(address string) (Balance, error){
    return c.decoder.DecodeBalance(c.Request("blockchain.address.get_balance", address))
}

//func (c *Client) AddressGetProof(address string) (AddressProof,error){
//    return c.decoder.DecodeAddressProof(c.Request("blockchain.address.get_proof", address))
//}

func (c *Client) AddressListUnspent(address string) ([]UnspentTransaction, error){
    return c.decoder.DecodeUnspent(c.Request("blockchain.address.listunspent", address))
}

//func (c *Client) GetAddress(utxo string) error{
//    response := c.Request("blockchain.utxo.get_address", utxo)
//    fmt.Println("response: ",response)
//    return nil
//}

func (c *Client) GetBlockHeader(height uint64) (BlockHeader, error){
    return c.decoder.DecodeBlockHeader(c.Request("blockchain.block.get_header", strconv.FormatUint(height, 10)))
}

func (c *Client) GetBlockChunk(chunk uint64) (string, error){
    return c.decoder.DecodeStringResult(c.Request("blockchain.block.get_chunk", strconv.FormatUint(chunk, 10)))
}

//func (c *Client) GetTransactionMerkle(txid string, height uint64) error{
//    response := c.Request("blockchain.transaction.get_merkle", txid, strconv.FormatUint(height, 10))
//    fmt.Println("response: ",response)
//    return nil
//}

func (c *Client) BroadcastTransaction(raw string) (string, error){
    return c.decoder.DecodeStringResult(c.Request("blockchain.transaction.broadcast", raw))

}
