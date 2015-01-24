package gostratum


import ("encoding/json")

type Encoder struct{}
type Decoder struct{}


func MakeEncoder() *Encoder{
    return &Encoder{}
}

func (e *Encoder) Encode(request *Request) (string, error){
    if(request.Params==nil){
        request.Params = make([]string, 0)
    }
    encoded,err:= json.Marshal(request)
    return string(encoded)+"\n",err
}


func MakeDecoder() *Decoder{
    return &Decoder{}
}

func (d *Decoder) Decode(msg string, response *Response) error{
    err:=json.Unmarshal([]byte(msg), &response)
    return err
}

func (d *Decoder) DecodeResult(response *Response, result interface{}) error{
    if response.Error !=nil {return response.Error}
    return d.DecodeData(response.Result, result);
}

func (d *Decoder) DecodeStringResult(response *Response) (string, error){
    var result string
    return result, d.DecodeResult(response, &result)
}


func (d *Decoder) DecodeData(response *json.RawMessage, result interface{}) error{
    if response==nil{
        return nil
    }
    return json.Unmarshal(*response, result)
}

func (d *Decoder) DecodePeers(response *json.RawMessage) ([]Peer, error){
    var peers []Peer
    var decoded [][3]interface{}
    err := d.DecodeData(response, &decoded)
    if err!=nil {
        return peers, err
    }
    for _,item := range decoded{
        var peer Peer
        var ok bool
        var info []interface{}

        peer.IP,ok = item[0].(string)
        if !ok {continue}

        peer.URL, ok = item[1].(string)
        if !ok {continue}

        info, ok = item[2].([]interface{})
        peer.Version, ok = info[0].(string)
        if !ok {continue}

        peers = append(peers, peer)
    }
    return peers, nil
}

func (d *Decoder) DecodeBlockHeaders(response *json.RawMessage) ([]BlockHeader, error){
    var blocks []BlockHeader
    err := d.DecodeData(response, blocks)
    return blocks, err
}

func (d *Decoder) DecodeInt(response *json.RawMessage) (int, error){
    var data int
    return data, d.DecodeData(response, &data)
}

func (d *Decoder) DecodeString(response *json.RawMessage) (string, error){
    var data string
    return data, d.DecodeData(response, &data)
}

func (d *Decoder) DecodeAddressTransactions(response *json.RawMessage) ([]AddressTransaction, error){
    var transactions []AddressTransaction
    return transactions, d.DecodeData(response, &transactions)
}

func (d *Decoder) DecodeAddressTransactionsResult(response *Response) ([]AddressTransaction, error){
    if response.Error != nil{
        return []AddressTransaction{}, response.Error
    }
    return d.DecodeAddressTransactions(response.Result)
}

func (d *Decoder) DecodeBalance(response *Response) (Balance, error){
    var balance Balance
    return balance, d.DecodeResult(response, &balance)
}

func (d *Decoder) DecodeUnspent(response *Response) ([]UnspentTransaction, error){
    var unspent []UnspentTransaction
    return unspent, d.DecodeResult(response, &unspent)
}

func (d *Decoder) DecodeBlockHeader(response *Response) (BlockHeader, error){
    var header BlockHeader
    return header, d.DecodeResult(response, &header)
}
