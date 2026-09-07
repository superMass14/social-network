export const JoinGroup = (groupId, sendMessage, readMessages) => {
  sendMessage({ type: "JoinGroup", groupId: groupId });
  readMessages();
};
