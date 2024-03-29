package youtube

import (
	"github.com/juju/errors"
	"regexp"
	"strings"
)

var videoRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`(?:v|embed|shorts|watch\?v)(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
	regexp.MustCompile(`([^"&?/=%]{11})`),
}

// ExtractVideoID extracts the videoID from the given string
func ExtractVideoID(videoID string) (string, error) {
	if strings.Contains(videoID, "youtu") || strings.ContainsAny(videoID, "\"?&/<%=") {
		for _, re := range videoRegexpList {
			if isMatch := re.MatchString(videoID); isMatch {
				subs := re.FindStringSubmatch(videoID)
				videoID = subs[1]
			}
		}
	}

	if strings.ContainsAny(videoID, "?&/<%=") {
		return "", errors.Annotate(ErrInvalidCharactersInVideoID, ErrInvalidCharactersInVideoID.Error())
	}

	if len(videoID) < 10 {
		return "", errors.Annotate(ErrVideoIDMinLength, ErrVideoIDMinLength.Error())
	}

	return videoID, nil
}
