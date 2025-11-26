// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var searchQuery = cli.Command{
	Name:  "query",
	Usage: "Get Spotify catalog information about albums, artists, playlists, tracks, shows,\nepisodes or audiobooks that match a keyword string. Audiobooks are only\navailable within the US, UK, Canada, Ireland, New Zealand and Australia markets.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "q",
			Usage: "Your search query.\n\nYou can narrow down your search using field filters. The available filters are `album`, `artist`, `track`, `year`, `upc`, `tag:hipster`, `tag:new`, `isrc`, and `genre`. Each field filter only applies to certain result types.\n\nThe `artist` and `year` filters can be used while searching albums, artists and tracks. You can filter on a single `year` or a range (e.g. 1955-1960).<br />\nThe `album` filter can be used while searching albums and tracks.<br />\nThe `genre` filter can be used while searching artists and tracks.<br />\nThe `isrc` and `track` filters can be used while searching tracks.<br />\nThe `upc`, `tag:new` and `tag:hipster` filters can only be used while searching albums. The `tag:new` filter will return albums released in the past two weeks and `tag:hipster` can be used to return only albums with the lowest 10% popularity.<br />\n",
		},
		&cli.StringSliceFlag{
			Name:  "type",
			Usage: "A comma-separated list of item types to search across. Search results include hits\nfrom all the specified item types. For example: `q=abacab&type=album,track` returns\nboth albums and tracks matching \"abacab\".\n",
		},
		&cli.StringFlag{
			Name:  "include-external",
			Usage: "If `include_external=audio` is specified it signals that the client can play externally hosted audio content, and marks\nthe content as playable in the response. By default externally hosted audio content is marked as unplayable in the response.\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of results to return in each item type.\n",
			Value: 20,
		},
		&cli.StringFlag{
			Name:  "market",
			Usage: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).\n",
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first result to return. Use\nwith limit to get the next page of search results.\n",
		},
	},
	Action:          handleSearchQuery,
	HideHelpCommand: true,
}

func handleSearchQuery(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.SearchQueryParams{
		Q:               cmd.Value("q").(string),
		Type:            cmd.Value("type").([]string),
		IncludeExternal: spotted.SearchQueryParamsIncludeExternal(cmd.Value("include-external").(string)),
		Market:          spotted.String(cmd.Value("market").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Search.Query(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("search query", json, format, transform)
}
