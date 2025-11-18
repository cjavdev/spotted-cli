// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var browseCategoriesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a single category used to tag items in Spotify (on, for example, the Spotify\nplayer’s “Browse” tab).",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "category-id",
			Usage: "The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) for the category.\n",
		},
		&cli.StringFlag{
			Name:  "locale",
			Usage: "The desired language, consisting of an [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), joined by an underscore. For example: `es_MX`, meaning &quot;Spanish (Mexico)&quot;. Provide this parameter if you want the category strings returned in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the specified language is not available, the category strings returned will be in the Spotify default language (American English)._\n",
		},
	},
	Action:          handleBrowseCategoriesRetrieve,
	HideHelpCommand: true,
}

var browseCategoriesList = cli.Command{
	Name:  "list",
	Usage: "Get a list of categories used to tag items in Spotify (on, for example, the\nSpotify player’s “Browse” tab).",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.StringFlag{
			Name:  "locale",
			Usage: "The desired language, consisting of an [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), joined by an underscore. For example: `es_MX`, meaning &quot;Spanish (Mexico)&quot;. Provide this parameter if you want the category strings returned in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the specified language is not available, the category strings returned will be in the Spotify default language (American English)._\n",
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
		},
	},
	Action:          handleBrowseCategoriesList,
	HideHelpCommand: true,
}

var browseCategoriesGetPlaylists = cli.Command{
	Name:  "get-playlists",
	Usage: "Get a list of Spotify playlists tagged with a particular category.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "category-id",
			Usage: "The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) for the category.\n",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.\n",
			Value: 20,
		},
		&cli.Int64Flag{
			Name:  "offset",
			Usage: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.\n",
		},
	},
	Action:          handleBrowseCategoriesGetPlaylists,
	HideHelpCommand: true,
}

func handleBrowseCategoriesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("category-id") && len(unusedArgs) > 0 {
		cmd.Set("category-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.BrowseCategoryGetParams{
		Locale: spotted.String(cmd.Value("locale").(string)),
	}
	var res []byte
	_, err := client.Browse.Categories.Get(
		ctx,
		cmd.Value("category-id").(string),
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
	return ShowJSON("browse:categories retrieve", json, format, transform)
}

func handleBrowseCategoriesList(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.BrowseCategoryListParams{
		Locale: spotted.String(cmd.Value("locale").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Browse.Categories.List(
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
	return ShowJSON("browse:categories list", json, format, transform)
}

func handleBrowseCategoriesGetPlaylists(ctx context.Context, cmd *cli.Command) error {
	client := spotted.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("category-id") && len(unusedArgs) > 0 {
		cmd.Set("category-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := spotted.BrowseCategoryGetPlaylistsParams{}
	if cmd.IsSet("limit") {
		params.Limit = spotted.Opt(cmd.Value("limit").(int64))
	}
	if cmd.IsSet("offset") {
		params.Offset = spotted.Opt(cmd.Value("offset").(int64))
	}
	var res []byte
	_, err := client.Browse.Categories.GetPlaylists(
		ctx,
		cmd.Value("category-id").(string),
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
	return ShowJSON("browse:categories get-playlists", json, format, transform)
}
