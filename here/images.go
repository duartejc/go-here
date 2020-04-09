package here

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/google/go-querystring/query"
)

// ImagesService provides for HERE Places api.
type ImagesService struct {
	httpClient *http.Client
	baseURL    string
}

// ImagesParams parameters for Images Service.
type ImagesParams struct {
	Waypoint0 string `url:"waypoint0"`
	Waypoint1 string `url:"waypoint1"`
	Poi0      string `url:"poix0"`
	Poi1      string `url:"poix1"`
	Poi2      string `url:"poix2"`
	Poi3      string `url:"poix3"`
	Poi4      string `url:"poix4"`
	Poi5      string `url:"poix5"`
	Poi6      string `url:"poix6"`
	Poithm    int    `url:"poithm"`
	// Poi1       string `url:"poi1"`
	// LineColor0 string `url:"lc0"`
	// LineColor1 string `url:"lc1"`
	// LineWidth0 string `url:"lw0"`
	// LineWidth1 string `url:"lw1"`
	// Resolution int    `url:"ppi"`
	// Width      int    `url:"w"`
	// Height     int    `url:"h"`
	APIKey string `url:"apikey"`
}

// newImagesService returns a new ImagesService.
func newImagesService(httpClient *http.Client, baseURL string) *ImagesService {
	return &ImagesService{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// Returns waypoints as a formatted string.
func createPoi(poi WaypointParams) string {
	return fmt.Sprintf("%f,%f;", poi.Coordinates[0], poi.Coordinates[1])
}

// CreateImagesParams creates images parameters struct.
func (s *ImagesService) CreateImagesParams(waypoints []WaypointParams, apiKey string) ImagesParams {

	stringWaypoint0 := createWaypoint(WaypointParams{Coordinates: waypoints[0].Coordinates})
	stringWaypoint1 := createWaypoint(WaypointParams{Coordinates: waypoints[len(waypoints)-1].Coordinates})

	imagesParams := ImagesParams{
		Waypoint0: stringWaypoint0,
		Waypoint1: stringWaypoint1,
		Poithm:    0,
		APIKey:    apiKey,
	}

	for i, waypoint := range waypoints[1 : len(waypoints)-1] {
		stringPoi := createPoi(waypoint)
		concatenated := "Poi" + strconv.Itoa(i)
		reflect.ValueOf(&imagesParams).Elem().FieldByName(concatenated).SetString(stringPoi)
	}

	return imagesParams
}

// Routing with given parameters.
func (s *ImagesService) Routing(params *ImagesParams) ([]byte, *http.Response, error) {
	apiError := new(APIError)

	v, _ := query.Values(params)

	buf := bytes.Buffer{}
	buf.WriteString(s.baseURL)
	buf.WriteString("routing?")
	buf.WriteString(v.Encode())

	reqURL := buf.String()

	fmt.Print(reqURL)
	resp, err := http.Get(reqURL)
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil, relevantError(err, *apiError)
}
