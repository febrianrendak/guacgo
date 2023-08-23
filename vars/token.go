package vars

/*
{
    "authToken": "AAC5EF59658532F8CA9BA0D05FF56707586C159FBA23BE99EC420E257C35567D",
    "username": "guacadmin",
    "dataSource": "postgresql",
    "availableDataSources": [
        "postgresql",
        "postgresql-shared"
    ]
}
*/

type TokenResp struct {
	AuthToken            string   `json:"authToken"`
	Username             string   `json:"username"`
	DataSource           string   `json:"datasource"`
	AvailableDataSources []string `json:"availableDataSources"`
}
