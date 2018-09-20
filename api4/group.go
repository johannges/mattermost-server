// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import "net/http"

func (api *API) InitGroup() {
	api.BaseRoutes.Groups.Handle("", api.ApiSessionRequired(createGroup)).Methods("POST")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}", api.ApiSessionRequiredTrustRequester(getGroup)).Methods("GET")
	api.BaseRoutes.Groups.Handle("", api.ApiSessionRequired(getGroups)).Methods("GET")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/patch", api.ApiSessionRequired(patchGroup)).Methods("PUT")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}", api.ApiSessionRequired(deleteGroup)).Methods("DELETE")

	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/members", api.ApiSessionRequired(createGroupMember)).Methods("POST")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/members/{user_id:[A-Za-z0-9]+}", api.ApiSessionRequired(deleteGroupMember)).Methods("DELETE")

	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/teams", api.ApiSessionRequired(createGroupTeam)).Methods("POST")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/teams", api.ApiSessionRequired(getGroupTeams)).Methods("GET")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/teams/{team_id:[A-Za-z0-9]+}", api.ApiSessionRequired(getGroupTeam)).Methods("GET")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/teams/{team_id:[A-Za-z0-9]+}/patch", api.ApiSessionRequired(patchGroupTeam)).Methods("PUT")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/teams/{team_id:[A-Za-z0-9]+}", api.ApiSessionRequired(deleteGroupTeam)).Methods("DELETE")

	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/channels", api.ApiSessionRequired(createGroupChannel)).Methods("POST")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/channels", api.ApiSessionRequired(getGroupChannels)).Methods("GET")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/channels/{channel_id:[A-Za-z0-9]+}", api.ApiSessionRequired(getGroupChannel)).Methods("GET")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/channels/{channel_id:[A-Za-z0-9]+}/patch", api.ApiSessionRequired(patchGroupChannel)).Methods("PUT")
	api.BaseRoutes.Groups.Handle("/{group_id:[A-Za-z0-9]+}/channels/{channel_id:[A-Za-z0-9]+}", api.ApiSessionRequired(deleteGroupChannel)).Methods("DELETE")
}

func createGroup(c *Context, w http.ResponseWriter, r *http.Request) {}
func getGroup(c *Context, w http.ResponseWriter, r *http.Request)    {}
func getGroups(c *Context, w http.ResponseWriter, r *http.Request)   {}
func patchGroup(c *Context, w http.ResponseWriter, r *http.Request)  {}
func deleteGroup(c *Context, w http.ResponseWriter, r *http.Request) {}

func createGroupMember(c *Context, w http.ResponseWriter, r *http.Request) {}
func deleteGroupMember(c *Context, w http.ResponseWriter, r *http.Request) {}

func createGroupTeam(c *Context, w http.ResponseWriter, r *http.Request) {}
func getGroupTeams(c *Context, w http.ResponseWriter, r *http.Request)   {}
func getGroupTeam(c *Context, w http.ResponseWriter, r *http.Request)    {}
func patchGroupTeam(c *Context, w http.ResponseWriter, r *http.Request)  {}
func deleteGroupTeam(c *Context, w http.ResponseWriter, r *http.Request) {}

func createGroupChannel(c *Context, w http.ResponseWriter, r *http.Request) {}
func getGroupChannels(c *Context, w http.ResponseWriter, r *http.Request)   {}
func getGroupChannel(c *Context, w http.ResponseWriter, r *http.Request)    {}
func patchGroupChannel(c *Context, w http.ResponseWriter, r *http.Request)  {}
func deleteGroupChannel(c *Context, w http.ResponseWriter, r *http.Request) {}
