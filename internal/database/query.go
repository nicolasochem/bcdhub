package database

func (d *db) GetOrCreateUser(u *User) error {
	return d.ORM.Where("login = ?", u.Login).FirstOrCreate(u).Error
}

func (d *db) GetUser(userID uint) (*User, error) {
	var user User

	if err := d.ORM.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *db) ListSubscriptions(userID uint) ([]Subscription, error) {
	var subs []Subscription

	if err := d.ORM.Where("user_id = ?", userID).Find(&subs).Error; err != nil {
		return nil, err
	}

	return subs, nil
}

func (d *db) CreateSubscription(s *Subscription) error {
	return d.ORM.Create(s).Error
}

func (d *db) DeleteSubscription(s *Subscription) error {
	return d.ORM.Delete(s).Error
}

func (d *db) GetSubscriptionRating(entityID string) (SubRating, error) {
	var s SubRating
	if err := d.ORM.Model(&Subscription{}).Where("entity_id = ?", entityID).Count(&s.Count).Error; err != nil {
		return s, err
	}

	if err := d.ORM.Raw(`
		SELECT users.login, users.avatar_url
		FROM subscriptions
		INNER JOIN users ON subscriptions.user_id=users.id
		WHERE entity_id = ?
		LIMIT 5;`, entityID).Scan(&s.Users).Error; err != nil {
		return s, err
	}

	return s, nil
}

func (d *db) Close() {
	d.ORM.Close()
}