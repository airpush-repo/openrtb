package openrtb

import "encoding/json"

// Source object describes the nature and behavior of the entity that is the source of the bid request upstream from the exchange.
type Source struct {
	FinalSaleDecision int             `json:"fd,omitempty"`     // Entity responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	TransactionID     string          `json:"tid,omitempty"`    // Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	PaymentChain      string          `json:"pchain,omitempty"` // Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol v1.0.
	SupplyChain       SupplyChain     `json:"schain,omitempty"` // This object represents both the links in the supply chain as well as an indicator whether or not the supply chain is complete.
	Ext               json.RawMessage `json:"ext,omitempty"`    // Placeholder for exchange-specific extensions to OpenRTB.
}

// SupplyChain object is composed primarily of a set of nodes where each node represents a specific entity that
// participates in the transacting of inventory. The entire chain of nodes from beginning to end represents all
// entities who are involved in the direct flow of payment for inventory.
type SupplyChain struct {
	Complete int               `json:"complete"` // Flag indicating whether the chain contains all nodes involved in the transaction leading back to the owner of the site, app or other medium of the inventory, where 0 = no, 1 = yes.
	Nodes    []SupplyChainNode `json:"nodes"`    // Array of SupplyChainNode objects in the order of the chain.
	Version  string            `json:"ver"`      // Version of the supply chain specification in use, in the format of “major.minor”.
	Ext      []byte            `json:"ext,omitempty"`
}

// SupplyChainNode object is associated with a SupplyChain object as an array of nodes. These nodes define the identity of
// an entity participating in the supply chain of a bid request
type SupplyChainNode struct {
	ASI    string `json:"asi"`              // The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that bidders connect to.
	SID    string `json:"sid"`              // The identifier associated with the seller or reseller account within the advertising system.
	RID    string `json:"rid,omitempty"`    // The OpenRTB RequestId of the request as issued by this seller.
	Name   string `json:"name,omitempty"`   // The name of the company (the legal entity) that is paid for inventory transacted under the given seller_id.
	Domain string `json:"domain,omitempty"` // The business domain name of the entity represented by this node.
	HP     int    `json:"hp,omitempty"`     // Indicates whether this node will be involved in the flow of payment for the inventory.
	Ext    []byte `json:"ext,omitempty"`
}
