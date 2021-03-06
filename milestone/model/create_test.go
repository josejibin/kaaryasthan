package milestone

import (
	"testing"

	"github.com/kaaryasthan/kaaryasthan/project/model"
	"github.com/kaaryasthan/kaaryasthan/test"
	"github.com/kaaryasthan/kaaryasthan/user/model"
)

func TestMilestoneCreate(t *testing.T) {
	t.Parallel()
	DB, conf := test.NewTestDB()
	defer test.ResetDB(DB, conf)

	usrDS := user.NewDatastore(DB)
	usr := &user.User{Username: "jack", Name: "Jack Wilber", Email: "jack@example.com", Password: "Secret@123"}
	if err := usrDS.Create(usr); err != nil {
		t.Fatal(err)
	}

	prjDS := project.NewDatastore(DB)
	prj := &project.Project{Name: "somename", Description: "Some description"}
	if err := prjDS.Create(usr, prj); err != nil {
		t.Error(err)
	}

	milDS := NewDatastore(DB)
	mil := &Milestone{Name: "somename", Description: "Some description", ProjectID: prj.ID}
	if err := milDS.Create(usr, mil); err != nil {
		t.Error(err)
	}
	if mil.ID <= 0 {
		t.Errorf("Data not inserted. ID: %#v", mil.ID)
	}
}
