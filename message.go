package gostratum

import ("encoding/json")


type Request struct{
    Method string         `json:"method"`
    Params []string       `json:"params"`
    ID     uint64         `json:"id"`
}


/*method, params is used in notifications*/

type Response struct{
    ID     uint64             `json:"id"`
    Result *json.RawMessage   `json:"result"`
    Error  error              `json:"error,string"`
    Method string             `json:"method"`
    Params *json.RawMessage   `json:"params"`
}



type BlockHeader struct{
    Nonce      uint64     `json:"nonce"`
    PrevHash   string     `json:"prev_block_hash"`
    Timestamp  uint64     `json:"timestamp"`
    Merkle     string     `json:"merkle_root"`
    Height     uint64     `json:"block_height"`
    UTXO       string     `json:"utxo_root"`
    Version    uint64     `json:"version"`
    Bits       uint64     `json:"bits"`
}

type Peer struct{
    IP      string
    URL     string
    Version string
}


type AddressTransaction struct{
    TxHash  string  `json:"tx_hash"`
    Height  uint64  `json:"height"`
}


type Balance struct{
    Confirmed    int64   `json:"confirmed"`
    Unconfirmed  int64   `json:"unconfirmed"`
}

//type AddressProof struct{
//    Address string 
//    Proof   string
//}


type UnspentTransaction struct{
    Hash    string  `json:"tx_hash"`
    Pos     uint64  `json:"tx_pos"`
    Value   uint64  `json:"value"`
    Height  uint64  `json:"height"`
}

type Transaction struct{

}
