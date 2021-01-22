package link_manager

import (
	om "github.com/pingguodeli573365/delinkcious/pkg/object_model"
)

type LinkStore interface {
	GetLinks(request om.GetLinksRequest) (om.GetLinksResult, error)
	AddLink(request om.AddLinkRequest) (*om.Link, error)
	UpdateLink(request om.UpdateLinkRequest) (*om.Link, error)
	DeleteLink(username string, url string) error
	SetLinkStatus(username, url string, status om.LinkStatus) error
}
