package link_manager

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	om "github.com/pingguodeli573365/delinkcious/pkg/object_model"
)

// User links are a map of url:TaggedLink
type UserLinks map[string]*om.Link

// Link store is a map of username:UserLinks
type inMemoryLinkStore struct {
	links map[string]UserLinks
}

func newInMemoryLinkStore() LinkStore {
	return &inMemoryLinkStore{map[string]UserLinks{}}
}

func (m *inMemoryLinkStore) GetLinks(request om.GetLinksRequest) (result om.GetLinksResult, err error) {
	result.Links = []om.Link{}
	userLinks := m.links[request.Username]
	if userLinks == nil {
		return
	}

	// Prepare complied regexes
	var urlRegex *regexp.Regexp
	var titleRegex *regexp.Regexp
	var descriptionRegex *regexp.Regexp
	if request.UrlRegex != "" {
		urlRegex, err = regexp.Compile(request.UrlRegex)
		if err != nil {
			return
		}
	}

	if request.TitleRegex != "" {
		titleRegex, err = regexp.Compile(request.UrlRegex)
		if err != nil {
			return
		}
	}

	if request.DescriptionRegex != "" {
		descriptionRegex, err = regexp.Compile(request.UrlRegex)
		if err != nil {
			return
		}
	}

	for _, link := range userLinks {
		// Check wach link against the regular expressions
		if urlRegex != nil && !urlRegex.MatchString(link.Url) {
			continue
		}

		if titleRegex != nil && !titleRegex.MatchString(link.Title) {
			continue
		}

		if descriptionRegex != nil && !descriptionRegex.MatchString(link.Description) {
			continue
		}

		// If there no tag was requested add link immediately and continue
		if request.Tag == "" {
			result.Links = append(result.Links, *link)
			continue
		}

		// Add link only if it has the request tag
		if link.Tags[request.Tag] {
			result.Links = append(result.Links, *link)
		}
	}

	return
}

func (m *inMemoryLinkStore) AddLink(request om.AddLinkRequest) (link *om.Link, err error) {
	if request.Url == "" {
		err = errors.New("URL can't be empty")
		return
	}

	if request.Username == "" {
		err = errors.New("user name can't be empty")
		return
	}

	userLinks := m.links[request.Username]
	if userLinks == nil {
		m.links[request.Username] = UserLinks{}
		userLinks = m.links[request.Username]
	} else {
		if userLinks[request.Url] != nil {
			msg := fmt.Sprintf("user %s already has a link for %s", request.Username, request.Url)
			err = errors.New(msg)
			return
		}
	}

	link = &om.Link{
		Url:         request.Url,
		Title:       request.Title,
		Description: request.Description,
		Status:      om.LinkStatusPending,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Tags:        request.Tags,
	}
	userLinks[request.Url] = link

	return
}

func (m *inMemoryLinkStore) UpdateLink(request om.UpdateLinkRequest) (link *om.Link, err error) {
	userLinks := m.links[request.Username]
	if userLinks == nil || userLinks[request.Url] == nil {
		msg := fmt.Sprintf("User %s doesn't have a link for %s", request.Username, request.Url)
		err = errors.New(msg)
		return
	}

	link = userLinks[request.Url]
	if request.Title != "" {
		link.Title = request.Title
	}

	if request.Description != "" {
		link.Description = request.Description
	}

	newTags := request.AddTags
	for t, _ := range link.Tags {
		if request.RemoveTags[t] {
			continue
		}

		newTags[t] = true
	}

	return
}

func (m *inMemoryLinkStore) DeleteLink(username string, url string) error {
	if url == "" {
		return errors.New("URL can't be empty")
	}

	if username == "" {
		return errors.New("User name can't be empty")
	}

	userLinks := m.links[username]
	if userLinks == nil || userLinks[url] == nil {
		msg := fmt.Sprintf("User %s doesn't have a link for %s", username, url)
		return errors.New(msg)
	}

	delete(m.links[username], url)
	return nil
}

func (m *inMemoryLinkStore) SetLinkStatus(username string, url string, status om.LinkStatus) error {
	if url == "" {
		return errors.New("URL can't be empty")
	}

	if username == "" {
		return errors.New("User name can't be empty")
	}

	userLinks := m.links[username]
	if userLinks == nil || userLinks[url] == nil {
		msg := fmt.Sprintf("User %s doesn't have a link for %s", username, url)
		return errors.New(msg)
	}

	userLinks[url].Status = status
	return nil
}
