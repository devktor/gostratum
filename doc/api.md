##gostratum API documentation

[func Connect(string)](#Connect)<br/>
[type Client](#Client)<br/>
[func (c *Client) SetTimeout(uint64)](#SetTimeout)<br/>
[func (c *Client) ServerVersion() (string, error)](#ServerVersion)<br/>
[func (c *Client) ServerBanner() (string, error)](#ServerBanner)<br/>
[func (c *Client) ServerDontationAddress() (string, error)](#ServerDontationAddress)<br/>
[func (c *Client) PeersSubscribe(callback func(\[\]Peer, error)) error](#PeersSubscribe)<br/>
[type Peer](#Peer)<br/>
[func (c *Client) BlockHeaderSubscribe(callback func(\[\]BlockHeader, error)) error](#BlockHeaderSubscribe)<br />
[func (c *Client) GetBlockHeader(height uint64) (BlockHeader, error)](#GetBlockHeader)<br />
[type BlockHeader](#BlockHeader)<br />
[func (c *Client) NumBlocksSubscribe(callback func(int, error)) error](#NumBlocksSubscribe)<br/>
[func (c *Client) AddressSubscribe(address string, callback func(string, string, error)) error](#AddressSubscribe)<br/>
[func (c *Client) AddressGetHistory(address string) (\[\]AddressTransaction, error)](#AddressGetHistory)<br />
[type AddressTransaction](#AddressTransaction)<br />
[func (c *Client) AddressGetBalance(address string) (Balance, error)](#AddressGetBalance)<br />
[type Balance](#Balance)<br />
[func (c *Client) AddressListUnspent(address string) (\[\]UnspentTransaction, error)](#AddressListUnspent)<br />
[type UnspentTransaction](#UnspentTransaction)<br />
[func (c *Client) GetBlockChunk(chunk uint64) (string, error)](#GetBlockChunk)<br />
[func (c *Client) BroadcastTransaction(raw string) (string, error)](#BroadcastTransaction)<br />
