package gostratum

import (
    "encoding/json"
)

func WrapPeersHandler(callback func([]Peer, error), decoder *Decoder) func(*json.RawMessage){
    return func(result* json.RawMessage){
        peers, err := decoder.DecodePeers(result)
        callback(peers, err)
    }
}

func WrapBlockHeadersHandler(callback func([]BlockHeader, error), decoder *Decoder) func(*json.RawMessage){
    return func(result *json.RawMessage){
        blocks, err := decoder.DecodeBlockHeaders(result)
        callback(blocks, err)
    }
}

func WrapAddressHandler(callback func(string, string, error), decoder *Decoder) func(*json.RawMessage){
    return func(result *json.RawMessage){
        var decoded [2]string
        err := decoder.DecodeData(result, &decoded)
        callback(decoded[0], decoded[1], err)
    }
}

func WrapNumBlocksHandler(callback func(int, error), decoder *Decoder) func(*json.RawMessage){
    return func(result *json.RawMessage){
        num, err := decoder.DecodeInt(result)
        callback(num, err)
    }
}

