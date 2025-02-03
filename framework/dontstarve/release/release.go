package release

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
)

const updateReleaseURL = "https://forums.kleientertainment.com/game-updates/dst"

const (
	BranchTest    = "Test"
	BranchRelease = "Release"
)

type Release struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	FullVersion string `json:"full_version"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Branch      string `json:"branch"`
	Pinned      bool   `json:"pinned"`
}

func New() *Client {
	return &Client{client: resty.New()}
}

func NewWithClient(client *resty.Client) *Client {
	return &Client{client: client}
}

type Client struct {
	client *resty.Client
}

// Pinned return the pinned latest release
func (c *Client) Pinned() (Release, error) {
	list, err := c.List(1)
	if err != nil {
		return Release{}, err
	}

	for _, release := range list {
		if release.Pinned {
			return release, nil
		}
	}

	return Release{}, err
}

// Latest return the latest release
func (c *Client) Latest(stable bool) (Release, error) {
	list, err := c.List(1)
	if err != nil {
		return Release{}, err
	}

	var maxRelease Release
	for _, release := range list {
		if !stable && release.Branch == BranchTest {
			continue
		}

		if release.ID > maxRelease.ID {
			maxRelease = release
		}
	}

	return maxRelease, err
}

// List returns the release list for the given page
func (c *Client) List(page int) ([]Release, error) {
	if page <= 0 {
		return nil, fmt.Errorf("invalid page number: %d", page)
	}

	resp, err := c.client.R().Get(fmt.Sprintf("%s/page/%d/", updateReleaseURL, page))
	if err != nil {
		return nil, err
	}
	defer resp.RawBody().Close()

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch updates: status code %d", resp.StatusCode())
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp.Body()))
	if err != nil {
		return nil, err
	}

	var releaseList []Release

	document.Find("a.cRelease").Each(func(i int, dom *goquery.Selection) {

		release := Release{}

		release.URL, _ = dom.Attr("href")

		id, _ := dom.Attr("data-releaseid")
		release.ID = cast.ToInt(id)

		if pinnedDom := dom.Find("span.ipsBadge_icon"); pinnedDom != nil {
			title, _ := pinnedDom.Attr("title")
			release.Pinned = title == "Pinned"
		}

		if typeDom := dom.Find("span.ipsType_large"); typeDom != nil {
			release.Type, _ = typeDom.Attr("title")
		}

		if head := dom.Find("h3.ipsType_sectionHead"); head != nil {
			release.Version = cast.ToInt(strings.TrimSpace(head.Contents().First().Text()))

			if badge := head.Find("span.ipsBadge"); badge != nil {
				release.Branch = strings.TrimSpace(badge.Contents().First().Text())
				release.Title, _ = badge.Attr("title")
			}
		}

		if dateDom := dom.Find(".ipsDataItem_meta"); dateDom != nil {
			dateStr := strings.TrimSpace(
				strings.TrimSuffix(
					strings.TrimPrefix(
						strings.TrimSpace(dateDom.Contents().First().Text()), "Released"), "..."))

			parse, err := time.Parse("01/02/06", dateStr)
			if err == nil {
				release.Date = parse.Format("2006-01-02")
			}
		}

		if release.Version != 0 && release.ID != 0 {
			release.FullVersion = fmt.Sprintf("%d-r%d", release.Version, release.ID)
		}

		releaseList = append(releaseList, release)
	})

	return releaseList, nil
}
