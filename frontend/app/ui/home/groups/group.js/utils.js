export const notGoing = (id_event, sendMessage, id_group) => {
  sendMessage({ type: "NotGoingEvent", groupId: id_group, event_id: id_event });
};
export const going = (id_event, sendMessage, id_group) => {
  sendMessage({ type: "GoingEvent", groupId: id_group, event_id: id_event });
};
