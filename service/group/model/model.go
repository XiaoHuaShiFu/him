package model

// Group 群
type Group struct {
	ID                           uint64 // 群编号
	Name                         string `gorm:"not null; size:20"`                  // 群名
	Icon                         string `gorm:"not null; size:200"`                 // 图标
	Members                      uint32 `gorm:"not null"`                           // 成员数
	Notice                       string `gorm:"not null; size:200"`                 // 群公告
	IsInviteJoinGroupNeedConfirm bool   `gorm:"not null; type:tinyint(1) unsigned"` // 是否邀请入群需要管理员或群主确认（默认不需要确认直接入群）,同时关闭扫码入群
	CreatorID                    uint64 `gorm:"not null; index"`                    // 创建者
	CreatedAt                    uint64 `gorm:"not null; index"`                    // 建立时间
	UpdatedAt                    uint64 `gorm:"not null; index"`                    // 修改时间
}

// GroupMember 群成员
type GroupMember struct {
	ID             uint64
	GroupID        uint64 `gorm:"not null; uniqueIndex:uk_group_id_member_id"` // 群编号
	MemberID       uint64 `gorm:"not null; uniqueIndex:uk_group_id_member_id"` // 成员编号
	Role           uint8  `gorm:"not null;  index"`                            // 角色
	GroupNickName  string `gorm:"not null; size:20"`                           // 群昵称
	IsDisturb      bool   `gorm:"not null; type:tinyint(1) unsigned"`          // 是否免打扰
	IsTop          bool   `gorm:"not null; type:tinyint(1) unsigned"`          // 是否置顶
	IsShowNickName bool   `gorm:"not null; type:tinyint(1) unsigned"`          // 是显示群成员昵称
	JoinTime       uint64 `gorm:"not null; index"`                             // 入群时间
	CreatedAt      uint64 `gorm:"not null; index"`                             // 建立时间
	UpdatedAt      uint64 `gorm:"not null; index"`                             // 修改时间
}

type JoinGroupEvent struct {
	ID                uint64
	GroupID           uint64 `gorm:"not null; index"`          // 群编号
	JoinGroupUserID   uint64 `gorm:"not null; index"`          // 入群用户编号
	JoinGroupUserMsg  string `gorm:"not null; size:50"`        // 入群用户消息
	InviterID         uint64 `gorm:"not null; index"`          // 邀请者编号
	InviterMsg        string `gorm:"not null; size:50"`        // 邀请人消息
	ManagerMsg        string `gorm:"not null; size:50"`        // 管理员消息
	Type              string `gorm:"not null; size:10; index"` // 入群事件类型
	ApplicationStatus string `gorm:"not null; size:10; index"` // 入群申请状态
	InviteStatus      string `gorm:"not null; size:10; index"` // 入群邀请状态
	CreatedAt         uint64 `gorm:"not null; index"`          // 创建时间
	UpdatedAt         uint64 `gorm:"not null; index"`          // 更新时间
}

// JoinGroupApplication 入群申请
type JoinGroupApplication struct {
	ID                uint64
	GroupID           uint64 `gorm:"not null; index"`          // 群编号
	ApplicantID       uint64 `gorm:"not null; index"`          // 申请者编号
	ApplicationMsg    string `gorm:"not null; size:50"`        // 申请消息
	ApplicationStatus string `gorm:"not null; size:10; index"` // 入群申请状态
	CreatedAt         uint64 `gorm:"not null; index"`          // 创建时间
	UpdatedAt         uint64 `gorm:"not null; index"`          // 更新时间
}
