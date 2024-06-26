package test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/feed3r/21Updater/src/model"
	"github.com/feed3r/21Updater/test/test_models"
	"github.com/feed3r/21Updater/test/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestParsePullRequest(t *testing.T) {

	logger := logrus.New()
	logger.Out = os.Stderr

	headers, err := utils.ReadHeaders(test_models.PR_HEADER)
	if err != nil {
		fmt.Println("Error in reading test string for headers: ", err)
		return
	}

	var prJson map[string]interface{}
	json.Unmarshal([]byte(test_models.PR_BODY), &prJson)

	var eventDesc = new(model.GHEventDescriptor)

	eventTranslator, err := utils.ReadDefaultLang()
	if err != nil {
		logger.Fatal("Error in reading the default language file: ", err)
	}

	eventDesc.Event = engine.ExtractEventFromHeader(&headers)
	require.Equal(t, "pull request", strings.ToLower(eventDesc.Event))

	engine.ParsePR(&headers, prJson, eventDesc, logger)
	require.Equal(t, test_models.PR_EXPECTED_TEXT, eventDesc.String(eventTranslator))

}
