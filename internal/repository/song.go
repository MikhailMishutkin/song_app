package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"song_app/internal/models"
	"strings"
)

// ...
func (r *Repo) CreateGroup(ctx context.Context, song *models.Song) (id int, err error) {
	r.log.Info("CreateGroup in repo started")

	const query = `
INSERT INTO groups (group_name) 
	VALUES ($1) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		song.GroupName,
	).Scan(&id)

	r.log.Debug("creategroup repo response values: ", "row", err, "id", id)
	if err != nil {
		return 0, fmt.Errorf("can't create group: %s", err)
	}
	return id, err
}

// ...
func (r *Repo) CreateSong(ctx context.Context, song *models.Song) (id int, err error) {
	r.log.Info("CreateSong in repo started")

	const query = `
INSERT INTO songs (song) 
	VALUES ($1) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		song.Song,
	).Scan(&id)

	r.log.Debug("createSong repo err: %v", "row", err, "id", id)
	if err != nil {
		return 0, fmt.Errorf("can't create song: %s", err)
	}
	return id, err
}

// ...
func (r *Repo) CreateSongUniqRec(ctx context.Context, groupId, songId int) (uniqId int, err error) {
	r.log.Info("CreateSongUniqRec in repo started")

	const query = `
INSERT INTO song_unique (group_id, song_id) 
	VALUES ($1, $2) RETURNING id
`
	err = r.DB.QueryRow(
		ctx,
		query,
		groupId,
		songId,
	).Scan(&uniqId)

	r.log.Debug("createSongUniq repo: %v", "row", err, "uniqid", uniqId)

	if err != nil {
		return 0, fmt.Errorf("can't create songUniqRec: %s", err)
	}
	return uniqId, err
}

// ...
func (r *Repo) AddDetails(ctx context.Context, song *models.Song) error {
	r.log.Info("AddDetails in repo started")

	const query = `
INSERT INTO details (uniq_id, release_date, text, link) 
	VALUES ($1, $2, $3, $4)
`
	tag, err := r.DB.Exec(
		ctx,
		query,
		song.Id,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)
	if err != nil {
		return fmt.Errorf("can't add details: %s", err)
	}
	r.log.Debug("do add details", "tag", tag)

	return err
}

// в теле PUT запроса передаются все данные песни как изменённые так и не изменённые, связку группа-песня нельзя изменить: по связке подгруженны обогащённые данные из API!!!
// также база спроектирована так, что изменение названия песни или группы повлечёт изменение во всех связанных таблицах:
// во всех песнях группы или во всех группах, которые песню с таким названием исполняли
func (r *Repo) UpdateSong(ctx context.Context, song *models.Song) error {
	r.log.Info("UpdateSong in repo started")

	const query = `
UPDATE details
SET release_date = $2, text = $3, link = $4 
WHERE uniq_id = $1
`
	_, err := r.DB.Exec(
		ctx,
		query,
		song.Id,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)
	if err != nil {
		return fmt.Errorf("can't update song: %s", err)
	}
	return err
}

// ...
func (r *Repo) DeleteSong(ctx context.Context, song *models.Song) error {
	r.log.Info("DeleteSong in repo started")

	_, err := r.DB.Exec(ctx, "DELETE FROM song_unique WHERE id = $1", song.Id)
	if err != nil {
		return fmt.Errorf("can't delete record: %v", err)
	}
	return err
}

// ...
func (r *Repo) GetSongText(ctx context.Context, fp *models.FiltAndPagin) (songText *models.Song, err error) {
	r.log.Info("GetSongText in repo started")

	var text string
	query := `SELECT text FROM details 
            WHERE uniq_id = $1
			`
	err = r.DB.QueryRow(ctx, query, fp.FilterMap["id"]).Scan(&text)

	if err != nil {
		return nil, fmt.Errorf("can't get song text: %s", err)
	}

	songText.Text = text

	return songText, err

}

// ...
func (r *Repo) GetAllSongs(ctx context.Context, fp *models.FiltAndPagin) (songs []*models.Song, err error) {
	r.log.Info("GetAllSongs in repo started")

	var rows pgx.Rows
	if len(fp.FilterMap) == 0 {
		rows, err = r.DB.Query(ctx, "SELECT song_unique.id, groups.group_name, songs.song, details.release_date, details.text, details.link FROM song_unique "+
			"INNER JOIN groups ON song_unique.group_id = groups.id "+
			"INNER JOIN songs ON song_unique.song_id = songs.id "+
			"INNER JOIN details ON song_unique.id = details.uniq_id ",
		)
	} else {
		rows, err = r.DB.Query(ctx, "SELECT song_unique.id, groups.group_name, songs.song, details.release_date, details.text, details.link FROM song_unique "+
			"INNER JOIN groups ON song_unique.group_id = groups.id "+
			"INNER JOIN songs ON song_unique.song_id = songs.id "+
			"INNER JOIN details ON song_unique.id = details.uniq_id "+
			" WHERE "+strings.Join(fp.Where, " AND "), fp.Values...)
	}

	r.log.Debug("pqx.Query result:", "row", rows)

	if err != nil {
		return nil, fmt.Errorf("something wrong with get songs info: %s", err)
	}

	for rows.Next() {
		song := &models.Song{}
		if err = rows.Scan(
			&song.Id,
			&song.GroupName,
			&song.Song,
			&song.ReleaseDate,
			&song.Text,
			&song.Link); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next then get songs with filter: %s", err)
		}

		songs = append(songs, song)
	}

	return songs, err
}

func (r *Repo) CheckExistGroup(ctx context.Context, song *models.Song) (id int, err error) {
	r.log.Info("CheckExistGroup in repo started")

	const query = `SELECT id FROM groups WHERE group_name = $1`

	err = r.DB.QueryRow(ctx, query, song.GroupName).Scan(&id)

	r.log.Debug("check group row existence", "row", err, "id", id)
	if err == pgx.ErrNoRows {

		return 0, nil
	}
	return id, err

}

func (r *Repo) CheckExistSong(ctx context.Context, song *models.Song) (id int, err error) {
	r.log.Info("CheckExistSong in repo started")

	const query = `SELECT id FROM songs where song = $1`

	err = r.DB.QueryRow(ctx, query, song.Song).Scan(&id)
	r.log.Debug("check song row existence", "row", err, "id", id)
	if err == pgx.ErrNoRows {

		return 0, nil
	}

	return id, err

}

func (r *Repo) CheckExistSongUniq(ctx context.Context, gr, s int) (id int, err error) {
	r.log.Info("CreateGroup in repo started")

	const query = `SELECT id FROM song_unique where group_id = $1 and song_id = $2`

	err = r.DB.QueryRow(ctx, query, gr, s).Scan(&id)
	r.log.Debug("check unique song row existence", "row", err, "id", id)
	if err == pgx.ErrNoRows {

		return 0, nil
	}

	return id, err

}
