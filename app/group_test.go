// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
)

func TestGetGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()
	group := th.CreateGroup()

	if _, err := th.App.GetGroup(group.Id); err != nil {
		t.Log(err)
		t.Fatal("Should get the group")
	}

	if _, err := th.App.GetGroup(model.NewId()); err == nil {
		t.Fatal("Should not have found a group")
	}
}

func TestGetGroupsPage(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()
	th.CreateGroup()
	th.CreateGroup()
	th.CreateGroup()

	groups, err := th.App.GetGroupsPage(1, 2)
	if err != nil {
		t.Log(err)
		t.Fatal("Should have groups")
	}

	if len(groups) < 1 {
		t.Fatal("Should have retrieved at least one group")
	}

	if groups, _ = th.App.GetGroupsPage(999, 1); len(groups) > 0 {
		t.Fatal("Should not have groups.")
	}
}

func TestCreateGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	group := &model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Type:        model.GroupTypeLdap,
	}

	if _, err := th.App.CreateGroup(group); err != nil {
		t.Log(err)
		t.Fatal("Should create a new group")
	}

	if _, err := th.App.CreateGroup(group); err == nil {
		t.Fatal("Should not create a new group - group already exist")
	}
}

func TestPatchGroup(t *testing.T) {}

func TestDeleteGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()
	group := th.CreateGroup()

	if _, err := th.App.DeleteGroup(group.Id); err != nil {
		t.Log(err)
		t.Fatal("Should delete the group")
	}

	if _, err := th.App.DeleteGroup(group.Id); err == nil {
		t.Fatal("Should not delete the group again - group already deleted")
	}
}
func TestCreateGroupMember(t *testing.T)  {}
func TestDeleteGroupMember(t *testing.T)  {}
func TestGetGroupTeam(t *testing.T)       {}
func TestCreateGroupTeam(t *testing.T)    {}
func TestDeleteGroupTeam(t *testing.T)    {}
func TestGetGroupChannel(t *testing.T)    {}
func TestCreateGroupChannel(t *testing.T) {}
func TestDeleteGroupChannel(t *testing.T) {}
