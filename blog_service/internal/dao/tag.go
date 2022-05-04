package dao

import (
	"goProjects/blog_service/internal/model"
	"goProjects/blog_service/pkg/app"
)

func (d *Dao)CountTag(name string,state uint8)(int64,error){
	tag:=model.Tag{Name: name,State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string,state uint8, page,pageSize int)([]*model.Tag,error){
	tag:=model.Tag{Name:name,State:state}
	pageOffset:= app.GetPageOffset(page,pageSize)
	return tag.List(d.engine,pageOffset,pageSize)
}

func (d *Dao)CreateTag(name string,state uint8,createBy string)error{
	tag:=model.Tag{
		Name:name,
		State:state,
		Model:&model.Model{CreatedBy: createBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao)UpdateTag(name string,state uint8,updatedBy string)error{
	tag:=model.Tag{
		Name:name,
		State:state,
		Model:&model.Model{ModifiedBy: updatedBy},
	}
	return tag.Update(d.engine)
}

func (d *Dao)DeleteTag(id uint32)error{
	tag:=model.Tag{
		Model:&model.Model{ID:id},
	}
	return tag.Delete(d.engine)
}