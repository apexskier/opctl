package engineclient

//go:generate counterfeiter -o ./fakeEngineClient.go --fake-name FakeEngineClient ./ EngineClient

import (
  "github.com/opspec-io/sdk-golang/pkg/models"
  "github.com/opspec-io/sdk-golang/pkg/engineprovider"
  "github.com/opspec-io/sdk-golang/util/format"
  "github.com/opspec-io/sdk-golang/util/http"
  netHttp "net/http"
)

type EngineClient interface {
  GetEventStream(
  ) (
  stream chan models.Event,
  err error,
  )

  KillOpRun(
  req models.KillOpRunReq,
  ) (
  err error,
  )

  StartOpRun(
  req models.StartOpRunReq,
  ) (
  opRunId string,
  err error,
  )
}

func New(
engineProvider engineprovider.EngineProvider,
) EngineClient {
  return &_engineClient{
    engineProvider:engineProvider,
    httpClient:netHttp.DefaultClient,
    jsonFormat: format.NewJsonFormat(),
    yamlFormat:format.NewYamlFormat(),
  }
}

type _engineClient struct {
  engineProvider engineprovider.EngineProvider
  httpClient http.Client
  jsonFormat format.Format
  yamlFormat format.Format
}
