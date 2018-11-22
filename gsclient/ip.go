package gsclient

import (
	"log"
)

type Ips struct {
	List map[string]IpProperties `json:"ips"`
}

type Ip struct {
	Properties IpProperties `json:"ip"`
}

type IpProperties struct {
	Name            string   `json:"name"`
	LocationCountry string   `json:"location_country"`
	LocationUuid    string   `json:"location_uuid"`
	ObjectUuid      string   `json:"object_uuid"`
	ReverseDns      string   `json:"reverse_dns"`
	Family          int      `json:"family"`
	Status          string   `json:"status"`
	CreateTime      string   `json:"create_time"`
	Failover        bool     `json:"failover"`
	ChangeTime      string   `json:"change_time"`
	LocationIata    string   `json:"location_iata"`
	LocationName    string   `json:"location_name"`
	Prefix          string   `json:"prefix"`
	Ip              string   `json:"ip"`
	Labels          []string `json:"labels"`
}

type IpCreateResponse struct {
	RequestUuid string `json:"request_uuid"`
	ObjectUuid  string `json:"object_uuid"`
	Prefix      string `json:"prefix"`
	Ip          string `json:"ip"`
}

type IpCreateRequest struct {
	Name         string        `json:"name,omitempty"`
	Family       int           `json:"family"`
	LocationUuid string        `json:"location_uuid"`
	Failover     bool          `json:"failover,omitempty"`
	ReverseDns   string        `json:"reverse_dns,omitempty"`
	Labels       []interface{} `json:"labels,omitempty"`
}

type IpUpdateRequest struct {
	Name       string        `json:"name,omitempty"`
	Failover   bool          `json:"failover"`
	ReverseDns string        `json:"reverse_dns,omitempty"`
	Labels     []interface{} `json:"labels,omitempty"`
}

func (c *Client) GetIp(id string) (*Ip, error) {
	r := Request{
		uri:    apiIpBase + "/" + id,
		method: "GET",
	}

	response := new(Ip)
	err := r.execute(*c, &response)

	log.Printf("Received ip: %v", response)

	return response, err
}

func (c *Client) GetIpList() (*Ips, error) {
	r := Request{
		uri:    apiIpBase,
		method: "GET",
	}

	response := new(Ips)
	err := r.execute(*c, &response)

	log.Printf("Received ips: %v", response)

	return response, err
}

func (c *Client) CreateIp(body IpCreateRequest) (*IpCreateResponse, error) {
	r := Request{
		uri:    apiIpBase,
		method: "POST",
		body:   body,
	}

	response := new(IpCreateResponse)
	err := r.execute(*c, &response)
	if err != nil {
		return nil, err
	}

	err = c.WaitForRequestCompletion(response.RequestUuid)

	return response, err
}

func (c *Client) DeleteIp(id string) error {
	r := Request{
		uri:    apiIpBase + "/" + id,
		method: "DELETE",
	}

	return r.execute(*c, nil)
}

func (c *Client) UpdateIp(id string, body IpUpdateRequest) error {
	r := Request{
		uri:    apiIpBase + "/" + id,
		method: "PATCH",
		body:   body,
	}

	return r.execute(*c, nil)
}

//Returns 0 if an error was encountered
func (c *Client) GetIpVersion(id string) int {
	ip, err := c.GetIp(id)
	if err != nil {
		return 0
	}
	return ip.Properties.Family
}
