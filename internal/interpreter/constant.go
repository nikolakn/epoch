package interpreter

const HELP string = `
help	
	document
		s | save                save
		q | exit | quit         exit
		open map                open map in browser if gps exist for event
		open url                open url in browser if url exist for event
		open html               export doc to html and open in browser
	add/delate
		a    | add                 add new event or epoch 
		del  | delate              delate of event or epoch 
		ae   | 'add event'         add new absolute event 
		are  | 'add rel event'     add new relative event 
		aep  | 'add epoch '        add new absolute epoch 
		arep | 'add rel epoch'     add new relative epoch 
	print
		p  | print              print timeline 
		pd | 'print des'        print description of event or epoch
		pr | 'print range'      print event between start and end date
		distance | dis          duration in years between start date of two event or epoch 
	edit
		r | rename | title      rename event or epoch 
		d | des                 change description of event or epoch 
		m | move                change start date of event or epoch 
		set                     set print options 
		show (flags, id, time, year only, gps, duration, description)
		hide (flags, id, time, year only, gps, duration, description)
		g | gps                 geo location; position for maps
		url | u                 url of event or epoch doc 
		importance | lvl        level of importance of event or epoch 
		type                    type of event or epoch 
	search
		search title | st       search by title
		search des   | sd       search by description
	`
