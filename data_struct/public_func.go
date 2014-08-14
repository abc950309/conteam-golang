package data_struct

func (this *Message) SetMessageID(vulae string) {
    this.MessageID = vulae
}

func (this *Message) GetMessageID() string {
    return this.MessageID
}

func (this *Message) SetFrom(vulae PublicPersonID) {
    this.From = vulae
}

func (this *Message) GetFrom() PublicPersonID {
    return this.From
}

func (this *Message) SetTo(vulae PublicPersonID) {
    this.To = vulae
}

func (this *Message) GetTo() PublicPersonID {
    return this.To
}

func (this *Message) SetContent(vulae string) {
    this.Content = vulae
}

func (this *Message) GetContent() string {
    return this.Content
}

func (this *Message) SetTimeStamp(vulae int64) {
    this.TimeStamp = vulae
}

func (this *Message) GetTimeStamp() int64 {
    return this.TimeStamp
}

func (this *PublicPersonID) SetUserID(vulae string) {
    this.UserID = vulae
}

func (this *PublicPersonID) GetUserID() string {
    return this.UserID
}

func (this *PublicPersonID) SetGroupID(vulae string) {
    this.GroupID = vulae
}

func (this *PublicPersonID) GetGroupID() string {
    return this.GroupID
}
