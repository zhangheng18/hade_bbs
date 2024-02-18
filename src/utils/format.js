import {dayjs}  from "element-plus";

export function formatDate(value) {
    console.log(value, 'value')

    if (value) {
        return dayjs(String(value)).format('YYYY-MM-DD HH:mm')
    }
    return value
}